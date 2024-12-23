package gosdk

import (
	"context"
	"log"

	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"google.golang.org/grpc"
)

// GetValidator retrieves a validator by address.
//
// @param addr the address of the validator
// @return the validator, or an error if retrieval fails
func (s *Server) GetValidator(addr string) (stakingtypes.Validator, error) {
	var result stakingtypes.Validator

	if err := s.KeepGrpcConn(); err != nil {
		log.Printf("error when keep grpc conn, endpoint: %v, err: %v", s.EndPoint, err.Error())
		return result, err
	}

	client := stakingtypes.NewQueryClient(s.Conn)

	req := &stakingtypes.QueryValidatorRequest{ValidatorAddr: addr}

	opt := make([]grpc.CallOption, 0)
	resp, err := client.Validator(context.Background(), req, opt...)
	if err != nil {
		log.Printf("could not query balances: %v", err)
		return result, err
	}

	return resp.Validator, nil
}

// GetValidatorList retrieves a list of validators with pagination.
//
// @param offset the offset for pagination
// @param pageSize the page size for pagination
// @return a list of validators, the total count, or an error if retrieval fails
func (s *Server) GetValidatorList(offset uint64, pageSize uint64) ([]stakingtypes.Validator, uint64, error) {
	if err := s.KeepGrpcConn(); err != nil {
		log.Printf("error when keep grpc conn, endpoint: %v, err: %v", s.EndPoint, err.Error())
		return nil, 0, err
	}

	client := stakingtypes.NewQueryClient(s.Conn)

	pagination := &query.PageRequest{
		Offset:     offset,
		Limit:      pageSize,
		CountTotal: true,
	}
	req := &stakingtypes.QueryValidatorsRequest{
		Status:     "",
		Pagination: pagination,
	}

	opt := make([]grpc.CallOption, 0)
	resp, err := client.Validators(context.Background(), req, opt...)
	if err != nil {
		log.Printf("could not query balances: %v", err)
		return nil, 0, err
	}

	return resp.Validators, resp.Pagination.Total, nil
}
