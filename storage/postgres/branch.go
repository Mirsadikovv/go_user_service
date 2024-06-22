package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	us "microservice/genproto/user_service"
	"microservice/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepoI {
	return &branchRepo{
		db: db,
	}
}

// Create implements storage.BranchRepoI.
func (u *branchRepo) Create(ctx context.Context, req *us.CreateBranch) (resp *us.Branch, err error) {
	resp = &us.Branch{}
	id := uuid.NewString()

	_, err = u.db.Exec(ctx, `
		INSERT INTO branch (
			id,
			phone,
			name,
			location,
			address,
			open_time,
			close_time,
			active
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		) `, id, req.Phone, req.Name, req.Location, req.Address, req.OpenTime, req.CloseTime, req.Active)

	if err != nil {
		log.Println("error while creating branch in storage", err)
		return nil, err
	}

	branch, err := u.GetByID(ctx, &us.BranchPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting branch by id", err)
		return nil, err
	}
	return branch, nil
}

// GetByID implements storage.BranchRepoI.
func (u *branchRepo) GetByID(ctx context.Context, req *us.BranchPrimaryKey) (*us.Branch, error) {
	resp := &us.Branch{}

	var (
		created_at sql.NullString
		updated_at sql.NullString
	)

	err := u.db.QueryRow(ctx, `
	        SELECT id,
	        phone,
			name,
			location,
			address,
			open_time,
			close_time,
			active,
	        created_at,
	        updated_at
	        FROM branch
	    WHERE id=$1`, req.Id).Scan(&resp.Id, &resp.Phone, &resp.Name, &resp.Location, &resp.Address,
		&resp.OpenTime, &resp.CloseTime, &resp.Active, &created_at, &updated_at)

	if err != nil {
		log.Println("error while getting branch by id", err)
		return nil, err
	}

	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String

	return resp, nil
}

// GetAll implements storage.BranchRepoI.h
func (u *branchRepo) GetList(ctx context.Context, req *us.GetListBranchRequest) (*us.GetListBranchResponse, error) {
	resp := &us.GetListBranchResponse{}

	var (
		filter     string
		created_at sql.NullString
		updated_at sql.NullString
		location   sql.NullString
		open_time  sql.NullString
		close_time sql.NullString
	)

	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = fmt.Sprintf(` AND (name ILIKE '%%%v%%' OR phone ILIKE '%%%v%%' OR email ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	rows, err := u.db.Query(ctx, `
		SELECT
			id,
			phone,
			name,
			location,
			address,
			open_time,
			close_time,
			active,
	        created_at,
	        updated_at
		FROM branch WHERE deleted_at is null
		 `+filter)

	if err != nil {
		log.Println("error while getting all branches", err)
		return nil, err
	}
	defer rows.Close()

	var count int64 = 0
	for rows.Next() {
		var branch us.Branch
		count++
		err = rows.Scan(&branch.Id, &branch.Phone, &branch.Name, &location, &branch.Address,
			&open_time, &close_time, &branch.Active, &created_at, &updated_at)

		if err != nil {
			log.Println("error while scanning branches", err)
			return nil, err
		}

		branch.Location = location.String
		branch.OpenTime = open_time.String
		branch.CloseTime = close_time.String
		branch.CreatedAt = created_at.String
		branch.UpdatedAt = updated_at.String

		resp.Branches = append(resp.Branches, &branch)
	}

	resp.Count = count

	return resp, nil
}

// Update implements storage.BranchRepoI.
func (u *branchRepo) Update(ctx context.Context, req *us.UpdateBranch) (resp *us.Branch, err error) {
	_, err = u.db.Exec(ctx, `
        UPDATE branch SET
            phone = $2,
            name = $3,
            location = $4,
            address = $5,
            open_time = $6,
            close_time = $7,
            active = $8,
            updated_at = NOW()
        WHERE id = $1`, req.Id, req.Phone, req.Name, req.Location, req.Address, req.OpenTime, req.CloseTime, req.Active)

	if err != nil {
		log.Println("error while updating branch in storage", err)
		return nil, err
	}

	branch, err := u.GetByID(ctx, &us.BranchPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting updated branch by id", err)
		return nil, err
	}

	return branch, nil
}

// Delete implements storage.BranchRepoI.
func (u *branchRepo) Delete(ctx context.Context, req *us.BranchPrimaryKey) (emptypb.Empty, error) {
	_, err := u.db.Exec(ctx, `
		UPDATE category SET
		active = false,
		deleted_at = NOW()
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting branch", err)
		return emptypb.Empty{}, err
	}

	return emptypb.Empty{}, nil
}
