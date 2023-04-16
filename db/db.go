package db

import (
	"context"
	"fmt"
	"gin-api/ent"
	"gin-api/ent/student"
	"gin-api/ent/user"
	"gin-api/internal/paging"
	"log"

	"github.com/gofrs/uuid"
)

func GetID() string {
	u, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	return u.String()
}

func GetAllStudent(ctx context.Context, client *ent.Client, paging *paging.Paging) ([]*ent.Student, error) {
	total, err := client.Student.Query().Count(ctx)
	if err != nil {
		log.Println(err)
	}

	paging.Total = total

	students, err := client.Student.Query().Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).All(ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func CreateStudent(ctx context.Context, client *ent.Client, name, age, school string) (*ent.Student, error) {
	s, err := client.Student.
		Create().
		SetID(GetID()).
		SetName(name).
		SetAge(age).
		SetSchool(school).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("student was created: ", s)
	return s, nil
}

func UpdateStudentByID(ctx context.Context, client *ent.Client, id, name, school, age string) error {
	if _, err := client.Student.UpdateOneID(id).SetName(name).SetAge(age).SetSchool(school).Save(ctx); err != nil {
		return err
	}
	return nil
}

func GetStudentByID(ctx context.Context, client *ent.Client, id string) (*ent.Student, error) {
	student, err := client.Student.
		Query().
		Where(student.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func DeleteStudentByID(ctx context.Context, client *ent.Client, id string) error {
	err := client.Student.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(ctx context.Context, client *ent.Client, name, username, email, password string) (*ent.User, error) {
	user, err := client.User.
		Create().
		SetName(name).
		SetUserName(username).
		SetPassword(password).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("student was created: ", user)
	return user, nil
}

func GetUserByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	user, err := client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
