package address

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example.com/rental/models"
	"github.com/gin-gonic/gin"
)

func SearchAddress(context *gin.Context){

	provided_address := context.Query("query")
    centerLat := context.Query("lat")
    centerLon := context.Query("lon")

	if provided_address == "" || centerLat == "" || centerLon == "" {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameters"})
        return
    }

    url := fmt.Sprintf("https://maps-api.pathao.io/v1/location/autocomplete/%s/%s/%s", centerLat, centerLon, provided_address)


	req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
        return
    }

	req.Header.Set("Authorization","BPARGUYMMSL4XFYZZ35CH546FORTV3EGEODHDZX44RNJGQH6MENG7UY24ERZDLVO55D66TQZ2GNZCQOCTUIDFAOFO6PDQUXSLWTM7SELOEPWX5MX6DG6LXISK2NPZLSSSYDUQT4X33AKKKLJM2RGLOCGAISUG6BWOWWSZYLOLJUWX2LCYPJYGACGQSCQ4RSZS22DV73FL7VFBRBODKL5D5TDQQXBLLOJQ2Z6Y2OCKD5LWRG2NR6D277HQQB2ISGLGHCQH7JN")
    
	client := &http.Client{}

	resp, err := client.Do(req)
	
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Request to Pathao API failed"})
        return
    }

    defer resp.Body.Close()



    body, err := io.ReadAll(resp.Body)

    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
        return
    }


    var SugAddress models.SuggestedAddress


    err = json.Unmarshal(body, &SugAddress); 
	
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Pathao response"})
        return
    }

    context.JSON(http.StatusOK, SugAddress.Results)


}