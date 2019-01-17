package main
import (
  "fmt"
  "io/ioutil"
  "encoding/xml"
  "os"
  "strings"
  "bytes"

)

type Project struct{
  XMLName xml.Name `xml:"project"`
	Dependencies   Dependencies  `xml:"dependencies"`
}
type Dependencies struct{
  XMLName xml.Name `xml:"dependencies"`
	Dependencies   []Dependency  `xml:"dependency"`
}
type Dependency struct{
  XMLName xml.Name `xml:"dependency"`
  GroupId string `xml:"groupId"`
  ArtifactId string `xml:"artifactId"`
  Version string `xml:"version"`
  Scope string `xml:"scope"`

}
func main(){
  fileName := os.Args[1]
  pomFile, err := os.Open(fileName)

  if err != nil{
    fmt.Println(err)
    os.Exit(2)
  }

  fmt.Println("Succcessfully Opened pom.xml : " , fileName)

  defer pomFile.Close()

  byteValue ,err := ioutil.ReadAll(pomFile)
  var deps Project
  err =xml.Unmarshal(byteValue, &deps)
  if err != nil{
    fmt.Println(err)
    os.Exit(2)
  }
  fmt.Println("Original Dependencies Count " ,len(deps.Dependencies.Dependencies))
  unique := SliceUniqMap(deps.Dependencies.Dependencies)
  fmt.Println("Unique Dependencies Count" ,len(unique))
  if (len(unique)== len(deps.Dependencies.Dependencies)){
    fmt.Println("No Duplicate Dependencies found for removal")
    os.Exit(2)
  }

  startTag := "<dependencies>"
  endTag := "</dependencies>"
  startIndex := strings.Index(string(byteValue),startTag)
  endIndex := strings.Index(string(byteValue),endTag)

  newContents :=(string(byteValue)[0:startIndex+len(startTag)]+ getUniqueDeps(unique) + string(byteValue)[endIndex:])


	err = ioutil.WriteFile(fileName, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}


}

func SliceUniqMap(s []Dependency) []Dependency {
	seen := make(map[Dependency]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

func getUniqueDeps(u []Dependency) string{
  var buffer bytes.Buffer

  for i:=0 ; i< len(u) ;i++{
    buffer.WriteString("\t <dependency> \n \t \t <groupdId>" +u[i].GroupId +"</groupdId> \n" )
    buffer.WriteString("\t \t <artifactId>" +u[i].ArtifactId +"</artifactId> \n" )
    if (len(u[i].Version) > 0) {
      buffer.WriteString("\t \t <version>" +u[i].Version +"</version> \n" )
    }
    if (len(u[i].Scope)> 0) {
      buffer.WriteString("\t \t <scope>" +u[i].Scope +"</scope> \n" )
    }
    buffer.WriteString("\t</dependency>\n")
  }
   return buffer.String()


}
