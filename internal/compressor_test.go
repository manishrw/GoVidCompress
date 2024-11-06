package compressor_test

import (
	compressor "GoVidCompressor/internal"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompressVideo(t *testing.T) {
	pathPrefix := "/mnt/c/Users/mwadhwani/.Personal"
	inputPath := pathPrefix + "/backup/i13pm/large-videos/IMG_3483.MOV"
	outputPath := pathPrefix + "/backup/i13pm/large-videos/IMG_3483_compressed.MOV"

	// Ensure the input file exists
	_, err := os.Stat(inputPath)
	require.NoError(t, err, "Input file does not exist")

	// Compress the video
	start := time.Now()
	err = compressor.CompressVideo(inputPath, outputPath)
	require.NoError(t, err, "Failed to compress video")
	duration := time.Since(start)

	// Ensure the output file exists
	_, err = os.Stat(outputPath)
	require.NoError(t, err, "Output file was not created")

	// Get the size of the input and output files
	inputFileInfo, err := os.Stat(inputPath)
	require.NoError(t, err, "Failed to get input file info")

	outputFileInfo, err := os.Stat(outputPath)
	require.NoError(t, err, "Failed to get output file info")

	// Check that the output file is smaller than the input file
	fmt.Println("Compression took: ", duration)
	fmt.Println("Original video size: ", inputFileInfo.Size()/1000000, "MB")
	fmt.Println("Compressed video size: ", outputFileInfo.Size()/1000000, "MB")
	fmt.Println("Compression ratio: ", inputFileInfo.Size()/outputFileInfo.Size(), "x")
	assert.True(t, outputFileInfo.Size() < inputFileInfo.Size(), "Compressed file is not smaller than the original file")

	// Clean up the output file after the test
	err = os.Remove(outputPath)
	require.NoError(t, err, "Failed to remove output file")
}
