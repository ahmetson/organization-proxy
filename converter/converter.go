package converter

import (
	"fmt"
	"github.com/ahmetson/common-lib/topic"
)

// SmartcontractFileName is the .json file name from the topic
func SmartcontractFileName(id topic.Topic) (string, error) {
	if !id.Has("org", "net", "name") {
		return "", fmt.Errorf("missing one of the parameters: 'org', 'net', 'name' in topic")
	}

	return fmt.Sprintf("%s.%s.%s.json", id.Organization, id.NetworkId, id.Name), nil
}

// ConfigurationFileName is the .json file name from the topic
func ConfigurationFileName(id topic.Id) (string, error) {
	parsed, err := id.Unmarshal()
	if !id.Has("org", "proj") {
		return "", fmt.Errorf("missing one of the parameters: 'org', 'net', 'name' in topic")
	}

	if err != nil {
		return "", fmt.Errorf("failed to umarshall id to topic: %w", err)
	}
	return fmt.Sprintf("%s.%s.json", parsed.Organization, parsed.Project), nil
}
