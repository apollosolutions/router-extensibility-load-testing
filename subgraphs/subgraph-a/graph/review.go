package graph

import "subgraph-a/graph/model"

var location1 = &model.Location{
	ID: "loc-1",
}

var location2 = &model.Location{
	ID: "loc-2",
}

var location3 = &model.Location{
	ID: "loc-3",
}

var location4 = &model.Location{
	ID: "loc-4",
}

var comment1 = "I would also like to say thank you to all your staff! I would gladly pay over 600 dollars for planet. Planet was worth a fortune to my company. After using planet my business skyrocketed!"
var comment2 = "It's really wonderful. We have no regrets! Keep up the excellent work."
var comment3 = "This is simply unbelievable! It's the perfect solution for our business. Really good. I don\"t always clop, but when I do, It's because of planet"
var comment4 = "Planet is exactly what our business has been lacking. It's incredible. If you want real marketing that works and effective implementation - planet\"s got you covered."
var comment5 = "Thanks planet! I was amazed at the quality of planet. Planet did exactly what you said it does."
var comment6 = "I would also like to say thank you to all your staff. I would gladly pay over 600 dollars for planet. Planet was worth a fortune to my company. After using planet my business skyrocketed!"
var comment7 = "It's really wonderful. We have no regrets! Keep up the excellent work."
var comment8 = "This is simply unbelievable! It's the perfect solution for our business. Really good. I don\"t always clop, but when I do, It's because of planet"
var comment9 = "Planet is exactly what our business has been lacking. It's incredible. If you want real marketing that works and effective implementation - planet's got you covered."
var comment10 = "Thanks planet! I was amazed at the quality of planet. Planet did exactly what you said it does."

var number2 = 2
var number3 = 3
var number4 = 4
var number5 = 5

var pointerReviews []*model.Review = []*model.Review{
	{
		ID:       "rev-1",
		Location: location1,
		Rating:   &number5,
		Comment:  &comment1,
	},
	{
		ID:       "rev-2",
		Location: location1,
		Rating:   &number5,
		Comment:  &comment2,
	},
	{
		ID:       "rev-3",
		Location: location1,
		Rating:   &number5,
		Comment:  &comment3,
	},
	{
		ID:       "rev-4",
		Location: location1,
		Rating:   &number2,
		Comment:  &comment4,
	},
	{
		ID:       "rev-5",
		Location: location2,
		Rating:   &number4,
		Comment:  &comment5,
	},
	{
		ID:       "rev-6",
		Location: location2,
		Rating:   &number3,
		Comment:  &comment6,
	},
	{
		ID:       "rev-7",
		Location: location2,
		Rating:   &number2,
		Comment:  &comment7,
	},
	{
		ID:       "rev-8",
		Location: location2,
		Rating:   &number5,
		Comment:  &comment8,
	},
	{
		ID:       "rev-9",
		Location: location3,
		Rating:   &number5,
		Comment:  &comment9,
	},
	{
		ID:       "rev-10",
		Location: location4,
		Rating:   &number5,
		Comment:  &comment10,
	},
}

func GetReviewsByLocationId(id string) ([]*model.Review, error) {
	var reviews []*model.Review
	for _, review := range pointerReviews {
		if review.Location.ID == id {
			reviews = append(reviews, review)
		}
	}
	return reviews, nil
}

func OverallRatingByLocationId(id string) (float64, error) {
	var reviews, _ = GetReviewsByLocationId(id)

	if len(reviews) == 0 {
		return 0, nil
	}

	var counter int = 0
	for _, review := range reviews {
		if review.Location.ID == id {
			counter += *review.Rating
		}
	}

	return float64(counter) / float64(len(reviews)), nil
}

func GetLocationsById(id string) (*model.Location, error) {
	var overallRating, err = OverallRatingByLocationId(id)

	if err != nil {
		return nil, err
	}

	var reviewsForLocation, err2 = GetReviewsByLocationId(id)

	if err2 != nil {
		return nil, err
	}

	return &model.Location{
		ID:                 id,
		OverallRating:      &overallRating,
		ReviewsForLocation: reviewsForLocation,
	}, nil
}
