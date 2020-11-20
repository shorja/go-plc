package plc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newCacheForTesting() (*Cache, DeviceFake) {
	devFake := DeviceFake(map[string]interface{}{})
	return NewCache(devFake), devFake
}

func TestNewCache(t *testing.T) {
	cache, _ := newCacheForTesting()
	assert.NotNil(t, cache, "NewCache should return a non-nil object")
}

func TestCachePassthroughInt(t *testing.T) {
	cache, deviceFake := newCacheForTesting()
	deviceFake[testTagName] = 7

	var actual int
	err := cache.ReadTag(testTagName, &actual)
	assert.NoError(t, err)
	assert.Equal(t, 7, actual)
}

func TestCachePassthroughTwiceInt(t *testing.T) {
	cache, deviceFake := newCacheForTesting()
	deviceFake[testTagName] = 7

	var unused int
	err := cache.ReadTag(testTagName, &unused)
	require.NoError(t, err)

	deviceFake[testTagName] = 85

	var actual int
	err = cache.ReadTag(testTagName, &actual)
	assert.NoError(t, err)
	assert.Equal(t, 85, actual)
}

func TestCacheInt(t *testing.T) {
	cache, deviceFake := newCacheForTesting()
	deviceFake[testTagName] = 7

	var unused int
	err := cache.ReadTag(testTagName, &unused)
	require.NoError(t, err)

	// Mess up both values to ensure we're not just getting a pointer
	unused++
	deviceFake[testTagName] = 13

	var actual int
	err = cache.ReadCachedTag(testTagName, &actual)
	assert.NoError(t, err)
	assert.Equal(t, 7, actual)
}

func TestCacheAfterUpdateInt(t *testing.T) {
	cache, deviceFake := newCacheForTesting()
	deviceFake[testTagName] = 7

	// Update the cache with a value we'll never look at
	var unused int
	err := cache.ReadTag(testTagName, &unused)
	require.NoError(t, err)

	// Set the value to 85 and read it, expecting this to end up in the cache
	deviceFake[testTagName] = 85
	err = cache.ReadTag(testTagName, &unused)
	require.NoError(t, err)

	var actual int
	err = cache.ReadCachedTag(testTagName, &actual)
	assert.NoError(t, err)
	assert.Equal(t, 85, actual)
}

const secondTestTagName = "secondTag"

func TestCacheSecondTag(t *testing.T) {
	cache, deviceFake := newCacheForTesting()
	deviceFake[testTagName] = 7
	deviceFake[secondTestTagName] = 99

	var unused int
	err := cache.ReadTag(testTagName, &unused)
	require.NoError(t, err)
	err = cache.ReadTag(secondTestTagName, &unused)
	require.NoError(t, err)

	// Ensure BOTH are read correctly
	var actual int
	err = cache.ReadCachedTag(testTagName, &actual)
	assert.NoError(t, err)
	assert.Equal(t, 7, actual)
	err = cache.ReadCachedTag(secondTestTagName, &actual)
	assert.NoError(t, err)
	assert.Equal(t, 99, actual)
}

func TestCacheReader(t *testing.T) {
	cache, deviceFake := newCacheForTesting()
	cacheReader := cache.CacheReader()
	deviceFake[testTagName] = 7

	var unused int
	err := cache.ReadTag(testTagName, &unused)
	require.NoError(t, err)

	// Mess up both values to ensure we're not just getting a pointer
	unused++
	deviceFake[testTagName] = 13

	var actual int
	err = cacheReader.ReadTag(testTagName, &actual)
	assert.NoError(t, err)
	assert.Equal(t, 7, actual)
}
