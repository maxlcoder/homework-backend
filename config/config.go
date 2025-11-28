package config

import (
	"bytes"
	"context"
	"log"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Config struct {
	Etcd struct {
		Endpoints []string
		Username  string
		Password  string
	}
	Database struct {
		Mysql struct {
			DNS string
		}
	}
	Kafka KafkaConfig
}

type KafkaConfig struct {
	Brokers []string
	Async   bool
	Topics  []string
}

var Conf *Config
var confMu sync.RWMutex

// Init 远程配置，从环境变量获取
func Init() {
	v := viper.New()
	// 1. 读取本地配置
	v.SetConfigName("config")
	v.AddConfigPath("./config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Println("read local config failed:", err)
	}

	// 重载配置
	load(v)

	// 2. 读取环境变量配置覆盖
	v.AutomaticEnv()
	Conf.Etcd.Username = v.GetString("ETCD_USERNAME")
	Conf.Etcd.Password = v.GetString("ETCD_PASSWORD")

	// 3. 监听本地文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("local config file changed:", e.Name)
		load(v)
	})
	// 4. 连接 etcd 获取远程配置
	watchRemoteConfig(v, false)
}

func load(v *viper.Viper) {
	confMu.Lock()
	defer confMu.Unlock()
	if Conf == nil {
		Conf = &Config{}
	}
	if err := v.Unmarshal(Conf); err != nil {
		log.Println("unmarshal config failed:", err)
		return
	}
	log.Println("config loaded successfully")
}

// 远程配置监听模式
func watchRemoteConfig(v *viper.Viper, isWatch bool) {
	etcdConfig := Conf.Etcd
	if len(etcdConfig.Endpoints) == 0 {
		return
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdConfig.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("connect to etcd failed:", err)
	}

	key := "/config/app.yml"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		log.Println("get etcd failed:", err)
	}
	if len(resp.Kvs) == 0 {
		log.Println("get etcd failed: key not found")
	}
	if err := v.MergeConfig(bytes.NewReader(resp.Kvs[0].Value)); err != nil {
		log.Println("read config failed:", err)
	}
	load(v)
	if isWatch {
		go func() {
			watchChan := cli.Watch(context.Background(), key)
			for wresp := range watchChan {
				for _, ev := range wresp.Events {
					if ev.Type == clientv3.EventTypePut {
						log.Println("config change event:", ev.Kv.Key, ev.Kv.Value)
						v.SetConfigType("yaml")
						if err := v.MergeConfig(bytes.NewReader(ev.Kv.Value)); err != nil {
							log.Println("merge config failed:", err)
							continue
						}
						load(v)
						log.Println("remote config update successfully")
					}
				}
			}
		}()
	}

}

// 远程配置轮询模式
func loadRemoteConfig(v *viper.Viper, isWatch bool) {
	etcdConfig := Conf.Etcd
	if len(etcdConfig.Endpoints) == 0 {
		return
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdConfig.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("connect to etcd failed:", err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	key := "/config/app.yml"
	resp, err := cli.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		log.Println("get etcd failed:", err)
	}
	if len(resp.Kvs) == 0 {
		log.Println("get etcd failed: key not found")
	}
	v.SetConfigType("yaml")
	if err := v.MergeConfig(bytes.NewReader(resp.Kvs[0].Value)); err != nil {
		log.Println("read config failed:", err)
	}

	load(v)

	if isWatch {
		// 远程热更新
		go func() {
			for {
				time.Sleep(10 * time.Second)
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				resp, err := cli.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSerializable())
				cancel()
				if err != nil {
					log.Println("get etcd failed:", err)
					continue
				}
				v.SetConfigType("yaml")
				if err := v.MergeConfig(bytes.NewReader(resp.Kvs[0].Value)); err != nil {
					log.Println("read config failed:", err)
				}
				load(v)
				log.Println("etcd remote config reloaded")
			}
		}()
	}
}

// 线程安全的获取配置
func GetConfig() *Config {
	confMu.RLock()
	defer confMu.RUnlock()
	return Conf
}
