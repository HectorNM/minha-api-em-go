package music

import ( 
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"errors"
)


type SongRepository struct {
	session *mgo.Session
}

const SongCollection = "songs"
var ErrDuplicatedSong = errors.New("Duplicated song")

func (r *SongRepository) FindById(id string) (*Song, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(SongCollection)
	query := bson.M{"_id": id}

	song := &Song{}

	err := collection.Find(query).One(song)
	return song, err
}

func (r *SongRepository) Create(p *Song) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(SongCollection)
	err := collection.Insert(p)
	mongoErr, ok := err.(*mgo.LastError)
	if ok && mongoErr.Code == 11000 {
		return ErrDuplicatedSong
	}
	return err
}

func (r *SongRepository) Update(p *Song) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(SongCollection)
	return collection.Update(bson.M{"_id": p.Id}, p)
}

func (r *SongRepository) Remove(id string) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(SongCollection)
	return collection.Remove(bson.M{"_id": id})
}

func (r *SongRepository) FindAllActive() ([]*Song, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(SongCollection)
	query := bson.M{"inative": false}

	documents := make([]*Song, 0)

	err := collection.Find(query).All(&documents)
	return documents, err
}

func NewSongRepository(session *mgo.Session) *SongRepository {
	return &SongRepository{session}
}