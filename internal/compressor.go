package compressor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// CompressVideo compresses the video at the given inputPath and saves the compressed video to outputPath.
func CompressVideo(inputPath string, outputPath string) error {
	// Check if the output file already exists
	if _, err := os.Stat(outputPath); err == nil {
		// Extract file directory, name, and extension
		_ = filepath.Dir(outputPath)
		ext := filepath.Ext(outputPath)
		base := outputPath[:len(outputPath)-len(ext)]

		// Append timestamp to the filename
		timestamp := time.Now().Format("20060102_150405")
		outputPath = fmt.Sprintf("%s_%s%s", base, timestamp, ext)
	}

	// ffmpeg command to compress video with a faster preset
	//cmd := exec.Command("ffmpeg", "-i", inputPath, "-vcodec", "libx265", "-preset", "fast", "-crf", "28", outputPath)
	cmd := exec.Command("ffmpeg",
		//"-hwaccel", "cuda", 	// Enable GPU acceleration (adjust as needed)
		"-i", inputPath,
		"-vcodec", "libx265",
		"-preset", "fast", // Set preset to "fast" for faster encoding
		"-threads", "8", // Use multiple threads
		"-crf", "24",
		outputPath,
	)

	/// Capture standard error to get ffmpeg’s output in case of errors
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to compress video: %w\nffmpeg error: %s", err, stderr.String())
	}

	return nil
}
