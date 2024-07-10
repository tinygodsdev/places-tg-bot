package bot

import (
	"context"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/data"
)

type fetchCitiesDataOutput struct {
	points  []data.Point
	sources []data.Source
	start   time.Time
}

type fetchCitiesDataInput struct {
	from time.Time
	to   time.Time
	city string
}

func (b *Bot) fetchCitiesData(input fetchCitiesDataInput) (*fetchCitiesDataOutput, error) {
	if input.from.IsZero() {
		input.from = time.Now().Add(-24 * time.Hour)
	}

	if input.to.IsZero() {
		input.to = time.Now()
	}

	var filterTags []data.Tag
	if input.city != "" {
		filterTags = append(filterTags, data.Tag{
			Label: TagCity,
			Value: input.city,
		})
	}

	start := time.Now()
	points, err := b.placesClient.GetPoints(
		context.TODO(),
		data.Filter{
			From: input.from,
			To:   input.to,
			Tags: filterTags,
		},
		data.Group{
			TagLabels: []string{TagCity},
		},
	)
	if err != nil {
		return nil, err
	}

	sources, err := b.placesClient.GetSources(context.TODO(), data.Filter{
		From: input.from,
		To:   input.to,
	})
	if err != nil {
		return nil, err
	}

	return &fetchCitiesDataOutput{
		points:  points,
		sources: sources,
		start:   start,
	}, nil
}
