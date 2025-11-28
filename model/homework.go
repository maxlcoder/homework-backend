package model

// School 学校
type School struct {
	BaseTenantModel
	Name    string `gorm:"size:60;not null;default:'';comment:学校名称"`
	Address string `gorm:"size:200;not null;default:'';comment:学校地址"`
}

type Teacher struct {
	BaseTenantModel
	Name string `gorm:"size:60;not null;default:'';comment:姓名"`
	Age  int8   `gorm:"comment:年龄"`
}

// SchoolTeacher 学校老师关系
type SchoolTeacher struct {
	BaseModel
	SchoolId  uint `gorm:"not null;default:0;comment:学校 ID"`
	TeacherId uint `gorm:"not null;default:0;comment:教师 ID"`
}

// Student 学生
type Student struct {
	BaseTenantModel
	Name     string `gorm:"size:60;not null;default:'';comment:学校名称"`
	SchoolId uint   `gorm:"not null;default:0;comment:学校 ID"`
	ClassId  uint   `gorm:"not null;default:0;comment:学校 ID"`
	Age      int8   `gorm:"comment:年龄"`
	Birthday string `gorm:"comment:生日"`
}

// Parent 家长
type Parent struct {
	BaseModel
	Name   string `gorm:"size:60;not null;default:'';comment:姓名"`
	Mobile string `gorm:"size:30;not null;default:'';comment:手机号"`
}

type StudentParent struct {
	BaseModel
	StudentId  uint `gorm:"not null;default:0;comment:学生 ID"`
	ParentId   uint `gorm:"not null;default:0;comment:家长 ID"`
	ParentType int8 `gorm:"not null;default:0;comment:家长类型 1:爸爸 2:妈妈 3:爷爷 4:奶奶 5..."`
}

// Class 班级
type Class struct {
	BaseTenantModel
	Name          string `gorm:"size:60;not null;default:'';comment:班级名称"`
	Number        string `gorm:"size:60;not null;default:'';comment:班级编号"`
	SchoolId      uint   `gorm:"not null;default:0;comment:学校 ID"`
	HeadTeacherId uint   `gorm:"not null;default:0;comment:班主任老师 ID"`
}

// Course 学科（课程）
type Course struct {
	BaseTenantModel
	Name      string `gorm:"size:60;not null;default:'';comment:课程名称"`
	Number    string `gorm:"size:60;not null;default:'';comment:课程编号"`
	SchoolId  uint   `gorm:"not null;default:0;comment:学校 ID"`
	ClassId   uint   `gorm:"not null;default:0;comment:班级 ID"`
	TeacherId uint   `gorm:"not null;default:0;comment:讲课老师 ID"`
}

// Assignment 作业  学校-班级-学科 作业
type Assignment struct {
	BaseTenantModel
	Title     string `gorm:"size:60;not null;default:'';comment:作业标题"`
	Number    string `gorm:"size:60;not null;default:'';comment:作业编号"`
	SchoolId  uint   `gorm:"not null;default:0;comment:学校 ID"`
	ClassId   uint   `gorm:"not null;default:0;comment:班级 ID"`
	TeacherId uint   `gorm:"not null;default:0;comment:老师 ID"`
	CourseId  uint   `gorm:"not null;default:0;comment:课程 ID"`
}

// AssignmentItem 作业项（结构化作业）
type AssignmentItem struct {
	BaseModel
	QuestionId uint   `gorm:"not null;default:0;comment:来源问题 ID"`
	Title      string `gorm:"size:200;not null;default:'';comment:标题"`
	Content    string `gorm:"type:text;comment:内容"`
	Remark     string `gorm:"size:2000;not null;default:'';comment:备注"`
	Score      int16  `gorm:"comment:分数"`
	Answer     string `gorm:"type:text;comment:答案"`
	Type       int8   `gorm:"comment:类型 1:填空题 2:选择题(单选) 3:选择题(多选) 4:对错题 5:解答题 6:画图题(附加)"`
}

// StudentAssignment 学生领取作业
type StudentAssignment struct {
	BaseTenantModel
	StudentId    uint  `gorm:"not null;default:0;comment:学生 ID"`
	AssignmentId uint  `gorm:"not null;default:0;comment:作业 ID"`
	Score        int16 `gorm:"comment:总分"`
}

// StudentAnswer 学生作业答案，答案可能有多种形式 一种文本形式（包含选择）一种文件形式
type StudentAnswer struct {
	BaseModel
	StudentAssigmentItemId uint   `gorm:"not null;default:0;comment:作业领取 ID"`
	AssigmentItemId        uint   `gorm:"not null;default:0;comment:作业项 ID"`
	StudentId              uint   `gorm:"not null;default:0;comment:学生 ID"`
	Answer                 string `gorm:"comment:答案"`
	AnswerFile             string `gorm:"comment:答案文件"`
	Score                  int16  `gorm:"comment:得分"`
	Revision               string `gorm:"comment:修订"`
	TeacherId              string `gorm:"comment:评分老师 ID"`
	TeacherComment         string `gorm:"comment:老师评语"`
	FinalScore             int16  `gorm:"comment:最终得分"`
}

// Question 问题
type Question struct {
	BaseModel
	Title   string `gorm:"size:200;not null;default:'';comment:标题"`
	Content string `gorm:"type:text;comment:内容"`
	Remark  string `gorm:"size:2000;not null;default:'';comment:备注"`
	Score   int16  `gorm:"comment:分数"`
	Answer  string `gorm:"type:text;comment:答案"`
	Type    int8   `gorm:"comment:类型 1:填空题 2:选择题(单选) 3:选择题(多选) 4:对错题 5:解答题 6:画图题(附加)"`
}
