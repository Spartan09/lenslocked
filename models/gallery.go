package models

import (
	"database/sql"
	"fmt"
	"github.com/Spartan09/lenslocked/errors"
	"path/filepath"
	"strings"
)

type Gallery struct {
	ID     int
	UserID int
	Title  string
}

type Image struct {
	Path string
}

type GalleryService struct {
	DB *sql.DB

	// ImagesDir is used to tell the GalleryService where to store and locate
	// images. If not set, the GalleryService will default to using the "images"
	// directory
	ImagesDir string
}

func (s *GalleryService) Create(title string, userID int) (*Gallery, error) {
	gallery := Gallery{
		Title:  title,
		UserID: userID,
	}
	row := s.DB.QueryRow(`
		INSERT INTO galleries (title, user_id)
		VALUES ($1, $2) RETURNING id;`, gallery.Title, gallery.UserID)
	err := row.Scan(&gallery.ID)
	if err != nil {
		return nil, fmt.Errorf("create gallery: %w", err)
	}
	return &gallery, nil
}

func (s *GalleryService) ByID(id int) (*Gallery, error) {
	gallery := Gallery{
		ID: id,
	}
	row := s.DB.QueryRow(`
		SELECT title, user_id
		FROM galleries
		WHERE id = $1;`, gallery.ID)
	err := row.Scan(&gallery.Title, &gallery.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("query gallery by id: %w", err)
	}
	return &gallery, nil
}

func (s *GalleryService) ByUserID(userID int) ([]Gallery, error) {
	rows, err := s.DB.Query(`
		SELECT id, title
		FROM galleries
		WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, fmt.Errorf("query galleries by user: %w", err)
	}
	var galleries []Gallery
	for rows.Next() {
		gallery := Gallery{
			UserID: userID,
		}
		err := rows.Scan(&gallery.ID, &gallery.Title)
		if err != nil {
			return nil, fmt.Errorf("query galleries by user: %w", err)
		}
		galleries = append(galleries, gallery)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("query galleries by user: %w", err)
	}
	return galleries, nil
}

func (s *GalleryService) Update(gallery *Gallery) error {
	_, err := s.DB.Exec(`
		UPDATE galleries
		SET title = $2
		WHERE id = $1;`, gallery.ID, gallery.Title)
	if err != nil {
		return fmt.Errorf("update gallery: %w", err)
	}
	return nil
}

func (s *GalleryService) Delete(id int) error {
	_, err := s.DB.Exec(`
		DELETE FROM galleries
		WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("delete gallery by id: %w", err)
	}
	return nil
}

func (s *GalleryService) galleryDir(id int) string {
	imagesDir := s.ImagesDir
	if imagesDir == "" {
		imagesDir = "images"
	}
	return filepath.Join(imagesDir, fmt.Sprintf("gallery-%d", id))
}

func (s *GalleryService) Images(galleryID int) ([]Image, error) {
	globPattern := filepath.Join(s.galleryDir(galleryID), "*")
	allFiles, err := filepath.Glob(globPattern)
	if err != nil {
		return nil, fmt.Errorf("retrieving gallery images: %w", err)
	}
	var images []Image
	for _, file := range allFiles {
		if hasExtension(file, s.extensions()) {
			images = append(images, Image{
				Path: file,
			})
		}
	}
	return images, nil
}

func (s *GalleryService) extensions() []string {
	return []string{".png", ".jpg", ".jpeg", ".gif"}
}

func hasExtension(file string, extensions []string) bool {
	for _, ext := range extensions {
		file = strings.ToLower(file)
		ext = strings.ToLower(ext)
		if filepath.Ext(file) == ext {
			return true
		}
	}
	return false
}
