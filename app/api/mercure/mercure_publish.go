package mercure

import (
	"bytes"
	"net/http"
)

func PublishUpdate(topic, data, jwt string) error {
    requestData := bytes.NewBufferString("topic=" + topic + "&data=" + data)
    request, err := http.NewRequest("POST", "http://localhost:9090/.well-known/mercure", requestData)
    if err != nil {
        return err
    }

    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Set("Authorization", "Bearer "+jwt)

    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    return nil
}



