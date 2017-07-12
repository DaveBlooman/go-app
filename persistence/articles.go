package persistence

import "github.com/DaveBlooman/go-app/models"

const (
	allArticles = `
    SELECT *
    FROM article.data
    ORDER BY publish_date
  `
)

func (persistence *Persistence) GetAllArticles() (*[]models.Mapping, error) {
	rows, err := persistence.db.Query(allArticles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Mapping

	for rows.Next() {
		var item models.Mapping
		err = rows.Scan(&item.Name, &item.Content, &item.PublishDate)
		if err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	return &results, nil
}
