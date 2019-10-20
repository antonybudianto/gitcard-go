package github

import (
	"context"
	pb "gogithub/protos"
	"log"
)

// GrpcServer is github grpc server
type GrpcServer struct{}

// FetchByUsername = implement from proto
func (s *GrpcServer) FetchByUsername(ctx context.Context, in *pb.GithubRequest) (*pb.GithubResponse, error) {
	log.Printf("[GithubGrpcServer] Received Username: %v", in.Username)
	data, err := FetchAllRepos(in.Username)
	if err != nil {
		log.Fatalf("failed to fetch github: %v", err)
	}

	return &pb.GithubResponse{
		Username:  in.Username,
		Starcount: int32(data.StarCount),
		Repocount: int32(data.RepoCount),
		Forkcount: int32(data.ForkCount),
		Langmap:   data.LanguageMap,
	}, nil
}
