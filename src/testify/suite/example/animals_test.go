package animals

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite functionality from testify
// - including a T() method which returns the current testing context
type AnimalAdapterTestSuite struct {
	suite.Suite
	HumanObj     *Human
	ClownfishObj *ClownFish
	SharkObj     *Shark
	BirdObj      *Bird
}

// Make sure that AnimalAdapterTestSuite struct data is set before each test
func (suite *AnimalAdapterTestSuite) SetupTest() {
	suite.HumanObj = new(Human)
	suite.HumanObj.Name = "Dr. Philip Sherman"
	suite.ClownfishObj = new(ClownFish)
	suite.ClownfishObj.Name = "Nemo"
	suite.SharkObj = new(Shark)
	suite.SharkObj.Name = "Bruce"
	suite.BirdObj = new(Bird)
	suite.BirdObj.Name = "Nigel"
}

// In order for 'go test' to run this suite, we need to create a normal test function and pass our suite to suite.Run
func TestAnimalAdapterTestSuite(t *testing.T) {
	suite.Run(t, new(AnimalAdapterTestSuite))
}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *AnimalAdapterTestSuite) TestAnimals() {

	assert.Equal(suite.T(), suite.HumanObj.Name, "Dr. Philip Sherman")
	suite.HumanObj.Move()

	assert.Equal(suite.T(), suite.ClownfishObj.Name, "Nemo")
	suite.ClownfishObj.Swim()

	assert.Equal(suite.T(), suite.SharkObj.Name, "Bruce")
	suite.SharkObj.Swim()

	assert.Equal(suite.T(), suite.BirdObj.Name, "Nigel")
	suite.BirdObj.Fly()
}

// Test the embedded adapter pattern and use of polymorphism
func (suite *AnimalAdapterTestSuite) TestMultiAnimals() {
	clownFishAdapter := ClownFishAdapter{suite.ClownfishObj}
	sharkAdapter := SharkAdapter{suite.SharkObj}

	birdAdapter := BirdAdapter{suite.BirdObj}
	assert.Equal(suite.T(), suite.BirdObj.Name, "Nigel")

	birdAdapter.Move()

	m := MultiAnimal{[]Animal{suite.HumanObj, &clownFishAdapter, &sharkAdapter, &birdAdapter}}
	m.Move()
}
