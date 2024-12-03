package transfer

import "fmt"

func (s *Service) GetUserID(email string) (string, error) {
	user, err := s.DirectoryClient.Users.Get(email).Do()
	if err != nil {
		return "", err
	}
	fmt.Println("got user", user.Id)

	return user.Id, nil
}
