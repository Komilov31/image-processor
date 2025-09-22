package service

import (
	"fmt"
	"log"

	"github.com/Komilov31/image-processor/internal/model"
	"github.com/google/uuid"
)

func (s *Service) GetImageStatus(id uuid.UUID) (*model.Image, error) {
	return s.storage.GetImageInfo(id)
}

func (s *Service) GetImageById(id uuid.UUID) (string, error) {
	imageInfo, err := s.storage.GetImageInfo(id)
	if err != nil {
		return "", err
	}

	if imageInfo.Status == "in progress" {
		return "", ErrNotProcessdYet
	}

	fileName := id.String() + "." + imageInfo.Format
	log.Println(fileName)
	if err := s.fileStorage.GetImage(fileName, processedDirName, "processed"); err != nil {
		return "", fmt.Errorf("could not get image from file storage: %w", err)
	}

	filePath := processedDirName + "/" + fileName
	return filePath, nil
}
