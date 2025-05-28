// Auto generated. DO NOT EDIT IT.
// Auto generated. DO NOT EDIT IT.
// Auto generated. DO NOT EDIT IT.

package models

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
)

// Override Singular Table name

type Tabler interface {
	TableName() string
  }
  
  func (Amphur) TableName() string {
    return "amphur"
}


func (Entrepreneur) TableName() string {
    return "entrepreneur"
}


func (Faculty) TableName() string {
    return "faculty"
}

func (Incharge) TableName() string {
    return "incharge"
}


func (Job) TableName() string {
    return "job"
}


func (Major) TableName() string {
    return "major"
}


func (Mooban) TableName() string {
    return "mooban"
}


func (Province) TableName() string {
    return "province"
}

func (Record) TableName() string {
    return "record"
}

func (Role) TableName() string {
    return "role"
}


func (Semester) TableName() string {
    return "semester"
}

func (Tambon) TableName() string {
    return "tambon"
}


func (Training) TableName() string {
    return "training"
}


func (User) TableName() string {
    return "user"
}

func (Plan) TableName() string {
    return "plain"
}

// Amphur [...]
type Amphur struct {
	ID         int64            `gorm:"autoIncrement:true;primaryKey;unique;column:id;type:int(11);not null" json:"id"`
	Value      string `gorm:"column:value;type:varchar(100);default:null" json:"value"`
	Code       int64  `gorm:"column:code;type:int(11);default:null" json:"code"`
	AmphurSeq  int64  `gorm:"index:outer_index;column:amphur_seq;type:int(11);default:null;default:0" json:"amphurSeq"`
	ProvinceID int64            `gorm:"primaryKey;index:fk_amphur_province1;column:province_id;type:int(11);not null" json:"provinceId"`
	Province Province `gorm:"foreignKey:ProvinceID" json:"province"`
}

// Entrepreneur [...]
type Entrepreneur struct {
	ID     int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	NameTh string `gorm:"column:name_th;type:varchar(100);default:null" json:"nameTh"`
	NameEn string `gorm:"column:name_en;type:varchar(100);default:null" json:"nameEn"`
	Tel    string `gorm:"column:tel;type:varchar(100);default:null" json:"tel"`
	Email  string `gorm:"column:email;type:varchar(100);default:null" json:"email"`
	Business string `gorm:"column:business;type:varchar(100);default:null" json:"business"`
	Employees int64 `gorm:"column:employees;type:int(11);default:null" json:"employees"`
	Manager string `gorm:"column:manager;type:varchar(100);default:null" json:"manager"`
	ManagerPosition string `gorm:"column:manager_position;type:varchar(100);default:null" json:"managerPosition"`
	ManagerDept string `gorm:"column:manager_dept;type:varchar(100);default:null" json:"managerDept"`
	Contact string `gorm:"column:contact;type:varchar(100);default:null" json:"contact"`
	ContactPosition string `gorm:"column:contact_position;type:varchar(100);default:null" json:"contactPosition"`
	ContactDept string `gorm:"column:contact_dept;type:varchar(100);default:null" json:"contactDept"`
	ContactTel string `gorm:"column:contact_tel;type:varchar(100);default:null" json:"contactTel"`
	ContactEmail string `gorm:"column:contact_email;type:varchar(100);default:null" json:"contactEmail"`
	Address     string `gorm:"column:address;type:varchar(100);default:null" json:"address"`
	Enable 		int64 `gorm:"column:enable;type:int(11);default:1" json:"enable"`
	MoobanID int64  `gorm:"index:fk_entrepreneur_mooban1_idx;column:mooban_id;type:int(11);default: null" json:"moobanId"`
	TambonID int64	`gorm:"index:fk_entrepreneur_tambon1_idx;column:tambon_id;type:int(11);default: null" json:"tambonId"`
	Mooban Mooban `gorm:"foreignKey:MoobanID" json:"mooban"` 
	Tambon Tambon `gorm:"foreignKey:TambonID" json:"tambon"` 
	Jobs  []Job
}

// Faculty [...]
type Faculty struct {
	ID        int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	FacultyEn string `gorm:"column:faculty_en;type:varchar(100);default:null" json:"facultyEn"`
	FacultyTh string `gorm:"column:faculty_th;type:varchar(100);default:null" json:"facultyTh"`
}

// Incharge [...]
type Incharge struct {
	ID             int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Fname          string `gorm:"column:fname;type:varchar(100);default:null" json:"fname"`
	Sname          string `gorm:"column:sname;type:varchar(100);default:null" json:"sname"`
	Position       string `gorm:"column:position;type:varchar(100);default:null" json:"position"`
	EntrepreneurID int64            `gorm:"index:fk_incharge_entrepreneur1_idx;column:entrepreneur_id;type:int(11);not null" json:"entrepreneurId"`
	Entrepreneur Entrepreneur `gorm:"foreignKey:EntrepreneurID" json:"entrepreneur"`
}

// Job [...]
type Job struct {
	ID             int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Name           string `gorm:"column:name;type:varchar(100);default:null" json:"name"`
	JobDes         string `gorm:"column:job_des;type:varchar(400);default:null" json:"jobDes"`
	EntrepreneurID int64            `gorm:"index:fk_incharge_entrepreneur1_idx;column:entrepreneur_id;type:int(11);not null" json:"entrepreneurId"`
	Entrepreneur   Entrepreneur `gorm:"foreignKey:EntrepreneurID" json:"entrepreneur"`
}

// Major [...]
type Major struct {
	ID        int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	MajorEn   string `gorm:"column:major_en;type:varchar(100);default:null" json:"majorEn"`
	MajorTh   string `gorm:"column:major_th;type:varchar(100);default:null" json:"majorTh"`
	Degree   string `gorm:"column:degree;type:varchar(100);default:null" json:"degree"`
	FacultyID int64            `gorm:"index:fk_major_faculty_idx;column:faculty_id;type:int(11);not null" json:"facultyId"`
	Faculty   Faculty  `gorm:"foreignKey:FacultyID" json:"faculty"`
}

// Mooban [...]
type Mooban struct {
	ID       int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Value    string `gorm:"column:value;type:varchar(100);default:null" json:"value"`
	Moo      int64            `gorm:"index:outer_index;column:moo;type:int(11);not null;default:0" json:"moo"`
	TambonID int64            `gorm:"index:fk_mooban_tambon1;column:tambon_id;type:int(11);not null" json:"tambonId"`
	Tambon   Tambon `gorm:"foreignKey:TambonID" json:"tambon"`
}

// Province [...]
type Province struct {
	ID    int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Value string `gorm:"column:value;type:varchar(100);default:null" json:"value"`
}

// Record [...]
type Record struct {
	ID         int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Week       int64  `gorm:"column:week;type:int(11);default:null" json:"week"`
	Startdate  datatypes.Date `gorm:"column:startdate;type:date;default:null" json:"startdate"`
	Enddate    datatypes.Date `gorm:"column:enddate;type:date;default:null" json:"enddate"`
	Starttime  sql.NullTime   `gorm:"column:starttime;type:time;default:null" json:"starttime"`
	Endtime    sql.NullTime   `gorm:"column:endtime;type:time;default:null" json:"endtime"`
	TrainingID int64            `gorm:"index:fk_record_training1_idx;column:training_id;type:int(11);not null" json:"trainingId"`
	Job        string `gorm:"column:job;type:mediumtext;default:null" json:"job"`
	Problem    string `gorm:"column:problem;type:mediumtext;default:null" json:"problem"`
	Fixed      string `gorm:"column:fixed;type:mediumtext;default:null" json:"fixed"`
	CourseFixed      string `gorm:"column:course_fixed;type:mediumtext;default:null" json:"course_fixed"`
	Exp        string `gorm:"column:exp;type:mediumtext;default:null" json:"exp"`
	Suggestion string `gorm:"column:suggestion;type:mediumtext;default:null" json:"suggestion"`
	Training   Training `gorm:"foreignKey:TrainingID" json:"training"`
}

// Role [...]
type Role struct {
	ID     int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Status string `gorm:"column:status;type:varchar(45);default:null" json:"status"`
	StatusEn string `gorm:"column:status_en;type:varchar(45);default:null" json:"statusEn"`
}

// Semester [...]
type Semester struct {
	ID       int64           `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Semester int64 `gorm:"column:semester;type:int(11);default:null" json:"semester"`
	Year     int64 `gorm:"column:year;type:int(11);default:null" json:"year"`
	IsCurrent     int64 `gorm:"column:is_current;type:int(11);default:0" json:"is_current"`
}

// Tambon [...]
type Tambon struct {
	ID        int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Value     string `gorm:"column:value;type:varchar(100);default:null" json:"value"`
	TambonSeq int64            `gorm:"index:outer_index;column:tambon_seq;type:int(11);not null;default:0" json:"tambonSeq"`
	AmphurID  int64            `gorm:"index:fk_tambon_amphur1;column:amphur_id;type:int(11);not null" json:"amphurId"`
	Amphur   Amphur `gorm:"foreignKey:AmphurID" json:"amphur"`
}

// Training [...]
type Training struct {
	UserID      int64            `gorm:"index:fk_user_has_semester_user1_idx;column:user_id;type:int(11);not null" json:"userId"`
	SemesterID  int64            `gorm:"index:fk_user_has_semester_semester1_idx;column:semester_id;type:int(11);not null" json:"semesterId"`
	ID          int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	JobID       int64            `gorm:"index:fk_training_job1_idx;column:job_id;type:int(11);not null" json:"jobId"`
	// for emergency Contact
	Address     string 			 `gorm:"column:address;type:varchar(100);default:null" json:"address"`
	MoobanID    int64            `gorm:"index:fk_training_mooban1_idx;null;column:mooban_id;type:int(11);default:null" json:"moobanId"`
	TambonID    int64            `gorm:"index:fk_training_tambon1_idx;column:tambon_id;type:int(11);default:null" json:"tambonId"`
	Tel    string `gorm:"column:tel;type:varchar(100);default:null" json:"tel"`
	Email  string `gorm:"column:email;type:varchar(100);default:null" json:"email"`

	Lat       float64            `gorm:"column:lat;type:double;default:null" json:"lat"`
	Long       float64            `gorm:"column:long;type:double;default:null" json:"long"`

	NameMentor     string 			 `gorm:"column:name_mentor;type:varchar(100);default:null" json:"nameMentor"`
	PositionMentor string 			 `gorm:"column:position_mentor;type:varchar(100);default:null" json:"positionMentor"`
	DeptMentor     string 			 `gorm:"column:dept_mentor;type:varchar(100);default:null" json:"deptMentor"`
	TelMentor      string 			 `gorm:"column:tel_mentor;type:varchar(100);default:null" json:"telMentor"`
	EmailMentor    string 			 `gorm:"column:email_mentor;type:varchar(100);default:null" json:"emailMentor"`

	JobPosition string `gorm:"column:job_position;type:varchar(100);default:null" json:"jobPosition"`
	JobDes      string `gorm:"column:job_des;type:varchar(400);default:null" json:"jobDes"`
	TimeMentor datatypes.Date `gorm:"column:time_mentor;type:date;default:null" json:"timeMentor"`


	InchargeID1 int64            `gorm:"index:fk_training_incharge1_idx;column:incharge_id1;type:int(11);default:null" json:"inchargeId1"`
	InchargeID2 int64  			`gorm:"index:fk_training_incharge2_idx;column:incharge_id2;type:int(11);default:null" json:"inchargeId2"`

	TeacherID1  int64            `gorm:"index:fk_training_user1_idx;column:teacher_id1;type:int(11);default:null" json:"teacherId1"`
	TeacherID2  int64            `gorm:"index:fk_training_user2_idx;column:teacher_id2;type:int(11);default:null" json:"teacherId2"`

	StartDate   time.Time `gorm:"column:startdate;type:date;default:null" json:"startDate"`
	EndDate     time.Time `gorm:"column:enddate;type:date;default:null" json:"endDate"`

	User   User `gorm:"foreignKey:UserID" json:"user"`
	Semester   Semester `gorm:"foreignKey:SemesterID" json:"semester"`
	Job   Job `gorm:"foreignKey:JobID" json:"job"`
	Mooban   Mooban `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:MoobanID" json:"mooban"`
	Tambon   Tambon `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TambonID" json:"tambon"`
	
	Incharge1   Incharge `gorm:"foreignKey:InchargeID1;references:ID" json:"incharge1"`
	Incharge2   Incharge `gorm:"foreignKey:InchargeID2;references:ID" json:"incharge2"`
	Teacher1   User `gorm:"foreignKey:TeacherID1" json:"teacher1"`
	Teacher2   User `gorm:"foreignKey:TeacherID2" json:"teacher2"`	
}

type Plan struct {
	ID         int64            `gorm:"autoIncrement:true;primaryKey;unique;column:id;type:int(11);not null" json:"id"`
	Month      int64 `gorm:"column:month;type:int(2);default:null" json:"month"`
	Topic      string `gorm:"column:topic;type:varchar(100);default:null" json:"topic"`
	TrainingID int64            `gorm:"index:fk_record_training1_idx;column:training_id;type:int(11);not null" json:"trainingId"`
	Training   Training `gorm:"foreignKey:TrainingID" json:"training"`
}


// User [...]
type User struct {
	ID        int64            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(11);not null" json:"id"`
	Username string  `gorm:"column:username;type:varchar(45);default:null" json:"username"`
	Title	  string `gorm:"column:title;type:varchar(45);default:null" json:"title"`
	Fname     string `gorm:"column:fname;type:varchar(45);default:null" json:"fname"`
	Sname     string `gorm:"column:sname;type:varchar(45);default:null" json:"sname"`
	MajorID   int64  `gorm:"index:fk_user_major1_idx;column:major_id;type:int" json:"majorId"`
	Picture   string        `gorm:"column:picture;type:varchar(100);default:null" json:"picture"`
	RoleID    int64            `gorm:"index:fk_user_role1_idx;column:role_id;type:int(11);not null;default:3" json:"roleId"`
	UpdatedAt sql.NullTime   `gorm:"column:updated_at;type:datetime;default:null;default:now()" json:"updatedAt"`
	CreatedAt sql.NullTime   `gorm:"column:created_at;type:datetime;default:null;default:now()" json:"createdAt"`
	TitleEn	  string `gorm:"column:title_en;type:varchar(45);default:null" json:"titleEn"`
	FnameEn   string `gorm:"column:fname_en;type:varchar(50);default:null" json:"fnameEn"`
	SnameEn   string `gorm:"column:sname_en;type:varchar(50);default:null" json:"snameEn"`
	IsAdmin int64 `gorm:"column:is_admin;type:int(11);default:0" json:"isAdmin"`
	Role	Role	`gorm:"foreignKey:RoleID" json:"role"`
	Major   Major  `gorm:"foreignKey:MajorID" json:"major"`
}
