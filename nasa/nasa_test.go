package nasa

import (
	"fmt"
	"testing"
)

func TestUnmarshall(t *testing.T) {
	s := "{" +
		"\"copyright\":\"CHART32 Team\"," +
		"\"date\":\"2021-07-22\"," +
		"\"explanation\":\"Point your telescope toward the high flying constellation Pegasus and you can find this expanse of Milky Way stars and distant galaxies. NGC 7814 is centered in the pretty field of view that would almost be covered by a full moon. NGC 7814 is sometimes called the Little Sombrero for its resemblance to the brighter more famous M104, the Sombrero Galaxy. Both Sombrero and Little Sombrero are spiral galaxies seen edge-on, and both have extensive halos and central bulges cut by a thin disk with thinner dust lanes in silhouette. In fact, NGC 7814 is some 40 million light-years away and an estimated 60,000 light-years across. That actually makes the Little Sombrero about the same physical size as its better known namesake, appearing smaller and fainter only because it is farther away. In this telescopic view from July 17, NGC 7814 is hosting a newly discovered supernova, dominant immediately to the left of the galaxy's core. Cataloged as SN 2021rhu, the stellar explosion has been identified as a Type Ia supernova, useful toward calibrating the distance scale of the universe.\"," +
		"\"hdurl\":\"https://apod.nasa.gov/apod/image/2107/NGC7814withSN2021rhuChart32.jpg\"," +
		"\"media_type\":\"image\"," +
		"\"service_version\":\"v1\"," +
		"\"title\":\"NGC 7814: Little Sombrero with Supernova\"," +
		"\"url\":\"https://apod.nasa.gov/apod/image/2107/NGC7814withSN2021rhuChart32_1024.jpg\"" +
		"}"

	var context = &NasaContext{}
	info, err := context.readDataInfo([]byte(s))
	if err != nil {
		t.Fatalf("Cannot parse")
	}

	if info.Url == "" {
		t.Fatalf("Cannot read")
	}
	fmt.Printf("%s\n", info.Url)
}

func TestGetImage(t *testing.T) {
	var context = &NasaContext{}
	file, err := context.getImage("https://apod.nasa.gov/apod/image/2107/NGC7814withSN2021rhuChart32_1024.jpg")
	if err != nil {
		t.Fatalf("Cannot parse")
	}

	fmt.Printf(string(file))
}
