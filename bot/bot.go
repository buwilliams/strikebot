package bot

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    dap "github.com/buwilliams/strikebot/adaptors"
)

func Strike(strikeMapPath string) {
    // list of adaptors
    adaptors := map[string]interface{}{
        "print": dap.Println,
    }

    // read json file
    j := readjson(strikeMapPath)
    steps := j["steps"].([]interface{})

    // loop through array of steps
    for i := range steps {
        obj := steps[i].(map[string]interface{})
        adaptor := obj["adaptor"].(string)
        options := obj["options"].(map[string]interface{})
        found := false

        for k, v := range adaptors {
            if adaptor == k {
                found = true
                v.(func(map[string]interface{}))(options)
            }
        }

        if found == false {
            fmt.Printf("Adaptor '%s' not found.\n", adaptor)
        }
    }
}

func readjson(strikeMapPath string) (map[string]interface{}){
    file, e := ioutil.ReadFile(strikeMapPath)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    //fmt.Printf("%s\n", string(file))

    var dat map[string]interface{}
    json.Unmarshal(file, &dat)
    //fmt.Printf("Results: %v\n", dat)

    return dat
}
