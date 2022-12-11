package repository

import (
	"github.com/VinayakBagaria/auth-micro-service/authentication/models"
	"github.com/VinayakBagaria/auth-micro-service/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const usersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	DeleteAll() error
}

type usersRepository struct {
	c *mgo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{conn.DB().C(usersCollection)}
}

func (r *usersRepository) Save(user *models.User) error {
	return r.c.Insert(user)
}

func (r *usersRepository) GetById(id string) (*models.User, error) {
	var user *models.User
	err := r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *usersRepository) GetByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.c.FindId(bson.M{"email": email}).One(&user)
	return user, err
}

func (r *usersRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := r.c.FindId(bson.M{}).All(&users)
	return users, err
}

func (r *usersRepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user)
}

func (r *usersRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

func (r *usersRepository) DeleteAll() error {
	return r.c.DropCollection()
}
