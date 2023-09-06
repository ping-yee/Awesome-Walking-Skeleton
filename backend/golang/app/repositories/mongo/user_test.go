package mongo

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/taimoor99/three-tier-golang/app/entities"
)

const email = "test@gmail.com"

type UserRepositoryTestSuite struct {
	suite.Suite
	UserRepository UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	m, err := GetSession(context.Background(), "test_db")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	m.Collection(entities.Users{}.UsersCollection()).Drop(context.Background())
	suite.UserRepository = NewUserRepository(m)
}

func TestOrderRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := entities.Users{}
	res, err := suite.UserRepository.CreateUser(context.Background(), user)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), res.ID)
}

func (suite *UserRepositoryTestSuite) TestFindUserById() {
	userBody := entities.Users{}
	user, err := suite.UserRepository.CreateUser(context.Background(), userBody)
	userFind, err := suite.UserRepository.FindUserById(context.Background(), user.ID.Hex())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), userFind.ID)
}

func (suite *UserRepositoryTestSuite) TestFindUserByEmail() {
	userBody := entities.Users{
		Email: email,
	}
	_, err := suite.UserRepository.CreateUser(context.Background(), userBody)
	userFind, err := suite.UserRepository.FindUserByEmail(context.Background(), email)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), userFind.ID)
	assert.Equal(suite.T(), userFind.Email, email)
}

func (suite *UserRepositoryTestSuite) TestGetAllUsers() {
	userBody := entities.Users{
		Email: email,
	}
	_, err := suite.UserRepository.CreateUser(context.Background(), userBody)
	userFind, err := suite.UserRepository.GetAllUsers(context.Background(), int64(10), int64(0))
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), len(userFind) == 1)
}
