package nbt

import (
	"encoding/json"
	"testing"
)

func TestJSONRoundtripSize(t *testing.T) {
	// Create a simple compound tag with a few elements
	compound := &TagCompound{
		baseTag: baseTag{
			tagType: BTagCompound,
			name:    "",
		},
		Value: []NBTTag{
			&TagInt{
				baseTag: baseTag{
					tagType: BTagInt,
					name:    "number",
				},
				Value: 42,
			},
			&TagString{
				baseTag: baseTag{
					tagType: BTagString,
					name:    "text",
				},
				Value: "hello",
			},
			&TagEnd{
				baseTag: baseTag{
					tagType: BTagEnd,
				},
			},
		},
	}

	// Serialize original to binary
	originalBytes, err := SerializeTag(compound, false)
	if err != nil {
		t.Fatalf("Failed to serialize original: %v", err)
	}
	t.Logf("Original binary size: %d bytes", len(originalBytes))

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(compound, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal to JSON: %v", err)
	}
	t.Logf("JSON size: %d bytes", len(jsonBytes))
	t.Logf("JSON:\n%s", string(jsonBytes))

	// Unmarshal from JSON
	var reconstructed TagCompound
	if err := json.Unmarshal(jsonBytes, &reconstructed); err != nil {
		t.Fatalf("Failed to unmarshal from JSON: %v", err)
	}

	// Serialize reconstructed to binary
	reconstructedBytes, err := SerializeTag(&reconstructed, false)
	if err != nil {
		t.Fatalf("Failed to serialize reconstructed: %v", err)
	}
	t.Logf("Reconstructed binary size: %d bytes", len(reconstructedBytes))

	// Compare sizes
	if len(originalBytes) != len(reconstructedBytes) {
		t.Errorf("Size mismatch: original=%d, reconstructed=%d (%.1fx)",
			len(originalBytes), len(reconstructedBytes),
			float64(len(reconstructedBytes))/float64(len(originalBytes)))
	}

	// Compare bytes
	if string(originalBytes) != string(reconstructedBytes) {
		t.Errorf("Binary content mismatch")
		t.Logf("Original bytes: %v", originalBytes)
		t.Logf("Reconstructed bytes: %v", reconstructedBytes)
	}
}

func TestJSONRoundtripWithList(t *testing.T) {
	// Create a compound with a list
	compound := &TagCompound{
		baseTag: baseTag{
			tagType: BTagCompound,
			name:    "root",
		},
		Value: []NBTTag{
			&TagList{
				baseTag: baseTag{
					tagType: BTagList,
					name:    "numbers",
				},
				ElementType: BTagInt,
				Value: []NBTTag{
					&TagInt{
						baseTag: baseTag{
							tagType: BTagInt,
						},
						Value: 10,
					},
					&TagInt{
						baseTag: baseTag{
							tagType: BTagInt,
						},
						Value: 20,
					},
					&TagInt{
						baseTag: baseTag{
							tagType: BTagInt,
						},
						Value: 30,
					},
				},
			},
			&TagEnd{
				baseTag: baseTag{
					tagType: BTagEnd,
				},
			},
		},
	}

	// Serialize original to binary
	originalBytes, err := SerializeTag(compound, false)
	if err != nil {
		t.Fatalf("Failed to serialize original: %v", err)
	}
	t.Logf("Original binary size: %d bytes", len(originalBytes))

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(compound, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal to JSON: %v", err)
	}
	t.Logf("JSON size: %d bytes", len(jsonBytes))
	t.Logf("JSON:\n%s", string(jsonBytes))

	// Unmarshal from JSON
	var reconstructed TagCompound
	if err := json.Unmarshal(jsonBytes, &reconstructed); err != nil {
		t.Fatalf("Failed to unmarshal from JSON: %v", err)
	}

	// Serialize reconstructed to binary
	reconstructedBytes, err := SerializeTag(&reconstructed, false)
	if err != nil {
		t.Fatalf("Failed to serialize reconstructed: %v", err)
	}
	t.Logf("Reconstructed binary size: %d bytes", len(reconstructedBytes))

	// Compare sizes
	if len(originalBytes) != len(reconstructedBytes) {
		t.Errorf("Size mismatch: original=%d, reconstructed=%d (%.1fx)",
			len(originalBytes), len(reconstructedBytes),
			float64(len(reconstructedBytes))/float64(len(originalBytes)))
	}

	// Compare bytes
	if string(originalBytes) != string(reconstructedBytes) {
		t.Errorf("Binary content mismatch")
		t.Logf("Original bytes: %v", originalBytes)
		t.Logf("Reconstructed bytes: %v", reconstructedBytes)
	}
}

func TestJSONRoundtripViaDeserialize(t *testing.T) {
	// This test simulates the workflow in main.go:
	// 1. Start with binary NBT data
	// 2. Deserialize to NBT structs (ParseNBT)
	// 3. Marshal to JSON
	// 4. Unmarshal from JSON back to NBT structs
	// 5. Serialize back to binary
	// 6. Compare sizes

	// Create original binary NBT data manually
	// TAG_Compound named "root" with:
	//   - TAG_Int named "number" = 42
	//   - TAG_String named "text" = "hello"
	//   - TAG_End
	originalBinary := []byte{
		0x0A,       // TAG_Compound
		0x00, 0x04, // name length = 4
		'r', 'o', 'o', 't', // name = "root"
		0x03,       // TAG_Int
		0x00, 0x06, // name length = 6
		'n', 'u', 'm', 'b', 'e', 'r', // name = "number"
		0x00, 0x00, 0x00, 0x2A, // value = 42
		0x08,       // TAG_String
		0x00, 0x04, // name length = 4
		't', 'e', 'x', 't', // name = "text"
		0x00, 0x05, // string length = 5
		'h', 'e', 'l', 'l', 'o', // value = "hello"
		0x00, // TAG_End
	}

	t.Logf("Original binary size: %d bytes", len(originalBinary))

	// Step 1: Deserialize binary to NBT structs (like ParseNBT does)
	parsedTag, err := ParseNBT(originalBinary, false)
	if err != nil {
		t.Fatalf("Failed to parse NBT: %v", err)
	}

	// Step 2: Marshal to JSON (like main.go does)
	jsonBytes, jsonErr := json.MarshalIndent(parsedTag, "", "  ")
	if jsonErr != nil {
		t.Fatalf("Failed to marshal to JSON: %v", jsonErr)
	}
	t.Logf("JSON size: %d bytes", len(jsonBytes))
	t.Logf("JSON:\n%s", string(jsonBytes))

	// Step 3: Unmarshal from JSON back to NBT structs (like serialize mode in main.go)
	var reconstructed TagCompound
	if unmarshalErr := json.Unmarshal(jsonBytes, &reconstructed); unmarshalErr != nil {
		t.Fatalf("Failed to unmarshal from JSON: %v", unmarshalErr)
	}

	// Step 4: Serialize back to binary (like SerializeTag does)
	reconstructedBinary, serializeErr := SerializeTag(&reconstructed, false)
	if serializeErr != nil {
		t.Fatalf("Failed to serialize reconstructed: %v", serializeErr)
	}
	t.Logf("Reconstructed binary size: %d bytes", len(reconstructedBinary))

	// Step 5: Compare
	if len(originalBinary) != len(reconstructedBinary) {
		t.Errorf("Size mismatch: original=%d, reconstructed=%d (%.1fx)",
			len(originalBinary), len(reconstructedBinary),
			float64(len(reconstructedBinary))/float64(len(originalBinary)))

		// Show hex dumps for debugging
		t.Logf("Original hex: % x", originalBinary)
		t.Logf("Reconstructed hex: % x", reconstructedBinary)
	}

	if string(originalBinary) != string(reconstructedBinary) {
		t.Errorf("Binary content mismatch")
	}
}

func TestJSONRoundtripNestedCompounds(t *testing.T) {
	// Create a compound with nested compounds
	// This tests if nested structures preserve their structure through JSON round-trip
	originalBinary := []byte{
		0x0A,       // TAG_Compound
		0x00, 0x04, // name length = 4
		'r', 'o', 'o', 't', // name = "root"

		// Inner compound
		0x0A,       // TAG_Compound
		0x00, 0x05, // name length = 5
		'i', 'n', 'n', 'e', 'r', // name = "inner"

		// Int inside inner compound
		0x03,       // TAG_Int
		0x00, 0x03, // name length = 3
		'n', 'u', 'm', // name = "num"
		0x00, 0x00, 0x00, 0x0A, // value = 10

		0x00, // TAG_End for inner compound

		0x00, // TAG_End for root compound
	}

	t.Logf("Original binary size: %d bytes", len(originalBinary))

	// Parse binary
	parsedTag, err := ParseNBT(originalBinary, false)
	if err != nil {
		t.Fatalf("Failed to parse NBT: %v", err)
	}

	// Marshal to JSON
	jsonBytes, jsonErr := json.MarshalIndent(parsedTag, "", "  ")
	if jsonErr != nil {
		t.Fatalf("Failed to marshal to JSON: %v", jsonErr)
	}
	t.Logf("JSON size: %d bytes", len(jsonBytes))
	t.Logf("JSON:\n%s", string(jsonBytes))

	// Unmarshal from JSON
	var reconstructed TagCompound
	if unmarshalErr := json.Unmarshal(jsonBytes, &reconstructed); unmarshalErr != nil {
		t.Fatalf("Failed to unmarshal from JSON: %v", unmarshalErr)
	}

	// Serialize back to binary
	reconstructedBinary, serializeErr := SerializeTag(&reconstructed, false)
	if serializeErr != nil {
		t.Fatalf("Failed to serialize reconstructed: %v", serializeErr)
	}
	t.Logf("Reconstructed binary size: %d bytes", len(reconstructedBinary))

	// Compare
	if len(originalBinary) != len(reconstructedBinary) {
		t.Errorf("Size mismatch: original=%d, reconstructed=%d (%.1fx)",
			len(originalBinary), len(reconstructedBinary),
			float64(len(reconstructedBinary))/float64(len(originalBinary)))

		t.Logf("Original hex: % x", originalBinary)
		t.Logf("Reconstructed hex: % x", reconstructedBinary)
	}

	if string(originalBinary) != string(reconstructedBinary) {
		t.Errorf("Binary content mismatch")
	}
}
