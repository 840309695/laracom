package repo

import (
	"github.com/jinzhu/gorm"
	pb "github.com/840309695/laracom/user-service/proto/user"
	"log"
	"os"
)

type Repository interface {
	Create(user *pb.User) error
	Get(id string) (*pb.User, error)
	GetByEmail(email string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	Update(user *pb.User) error
}

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	user := &pb.User{}
	user.Id = id
	if err := repo.Db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.Db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {

	var users []*pb.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	file := "./" +"message"+ ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.Println(users) // log 还是可以作为输出的前缀
	return users, nil
}

func (repo *UserRepository) Update(user *pb.User) error {
	if err := repo.Db.Save(user).Error; err != nil {
		return err
	}
	return nil
}