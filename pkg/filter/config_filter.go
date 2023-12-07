package filter

import (
	"log"
	"strconv"

	"groupie-tracker/internal/models"
)

func FilterConfig(Bands []models.Artists) (models.Filters, error) {
	min_creation_date := 2100
	max_creation_date := 0

	min_album_date := 2100
	max_album_date := 0

	min_members := 100
	max_members := 0

	for _, v := range Bands {
		if v.CreationDate < min_creation_date {
			min_creation_date = v.CreationDate
		}

		if v.CreationDate > max_creation_date {
			max_creation_date = v.CreationDate
		}

		year, err := strconv.Atoi(v.FirstAlbum[6:])
		if err != nil {
			log.Fatal("Error on converting date string to date int")
			return models.Filters{}, err
		}

		if year < min_album_date {
			min_album_date = year
		}

		if year > max_album_date {
			max_album_date = year
		}

		if len(v.Members) < min_members {
			min_members = len(v.Members)
		}

		if len(v.Members) > max_members {
			max_members = len(v.Members)
		}
	}

	return models.Filters{
		MinCreationDate:   min_creation_date,
		MaxCreationDate:   max_creation_date,
		MinFirstAlbumDate: min_album_date,
		MaxFirstAlbumDate: max_album_date,
		MinMembers:        min_members,
		MaxMembers:        max_members,
	}, nil
}
