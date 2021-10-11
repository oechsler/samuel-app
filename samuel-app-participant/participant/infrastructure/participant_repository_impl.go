package infrastructure

import (
	"github.com/cozy-hosting/clerk"
	"github.com/oechsler/samuel-app-participant/participant/domain"
	"os"
)

type ParticipantRepositoryImpl struct {
	connection clerk.Connection
	collection clerk.Collection
}

func NewParticipantRepositoryImpl(connection clerk.Connection) domain.ParticipantRepository {
	databaseName := os.Getenv("MONGO_DB_DATABASE")
	collection := clerk.NewDatabase(databaseName).GetCollection("participants")

	return &ParticipantRepositoryImpl{connection: connection, collection: collection}
}

func (p ParticipantRepositoryImpl) StoreParticipant(participant *domain.Participant) error {
	storeCommand := clerk.NewMongoUpdateCommand(p.collection, participant).
		Where("_id", participant.Id).
		WithUpsert()

	return p.connection.SendCommand(storeCommand)
}

func (p ParticipantRepositoryImpl) GetParticipantById(participantId string) (*domain.Participant, error) {
	getByIdQuery := clerk.NewMongoSingleQuery(p.collection).Where("_id", participantId)

	iterator, err := p.connection.SendQuery(getByIdQuery)
	if err != nil {
		return nil, err
	}

	var participant *domain.Participant
	if err := iterator.Single(&participant); err != nil {
		return nil, err
	}

	return participant, nil
}

func (p ParticipantRepositoryImpl) ListParticipants() ([]domain.Participant, error) {
	listQuery := clerk.NewMongoListQuery(p.collection)

	iterator, err := p.connection.SendQuery(listQuery)
	if err != nil {
		return nil, err
	}

	var participants []domain.Participant
	for iterator.Next() {
		var participant domain.Participant
		if err := iterator.Decode(&participant); err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}

	return participants, nil
}
