# go-compass
Compute compass directions / cardinal directions


# Install
````
go get github.com/flopp/go-compass
````

# Use
````
import "github.com/flopp/go-compass"

func main() {
    // Resolution4: N/E/S/W
    dir := compass.GetDirection(13.0, compass.Resolution4)
    // => compass.N

    // Resolution8: N/NE/E/SE/S/SW/W/NW
    dir = compass.GetDirection(13.0, compass.Resolution4)
    // => compass.N

    // Resolution16: N/NNE/NE/ENE/E/ESE/SE/SSE/S/SSW/SW/WSW/W/WNW/NW/NNW
    dir = compass.GetDirection(13.0, compass.Resolution16)
    // => compass.NNE
}
````
