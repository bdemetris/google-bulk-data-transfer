package transfer

import (
	"fmt"

	admin "google.golang.org/api/admin/datatransfer/v1"
)

func (s *Service) RequestTransfer(src, dst string) error {
	srcId, err := s.GetUserID(src)
	if err != nil {
		return err
	}

	dstId, err := s.GetUserID(dst)
	if err != nil {
		return err
	}

	fmt.Println("source user", src, "has id", srcId)
	fmt.Println("destination user", dst, "has id", dstId)

	config := admin.DataTransfer{
		NewOwnerUserId: dstId,
		OldOwnerUserId: srcId,
		ApplicationDataTransfers: []*admin.ApplicationDataTransfer{
			{
				ApplicationId: 55656082996,
				ApplicationTransferParams: []*admin.ApplicationTransferParam{
					{
						Key:   "PRIVACY_LEVEL",
						Value: []string{"SHARED", "PRIVATE"},
					},
				},
			},
		},
	}

	resp, err := s.TransferClient.Transfers.Insert(&config).Do()
	if err != nil {
		return err
	}

	fmt.Println(resp.OverallTransferStatusCode)
	return nil
}
