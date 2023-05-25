import { markdownTable } from "markdown-table";
import fs from 'fs'
import path from 'path'
import { versions } from "process";

const regex = /results_(.*).json/
// headers for MD table
const tableHeaders = ['Type', 'Min (ms)', 'Mean (ms)', 'p50 (ms)', 'p90 (ms)', 'p95 (ms)', 'p99 (ms)', 'Max (ms)']
let testResults = {}

// get the test results from the directory
const getTestResults = async () => {
    try {
        // the script runs relative to invocation, and the taskfile invokes from within this directory- hence the ..
        let testFolder = '../tests'
        let files = await fs.promises.readdir(testFolder)
        for await (let file of files) {
            let stat = await fs.promises.stat(path.join(testFolder, file))
            if (stat.isDirectory()) {
                // convert the results into the format necessary for the markdownTable function
                await convertResults(path.join(testFolder, file), file)
            }
        }
        // now that we have results, convert into MD and log out to be piped to the results.md file
        for (let test in testResults) {
            let tr = testResults[test]
            // check if the title/description exist before using; otherwise just skip and default to the raw name
            tr.title ? console.log(`### ${testResults[test].title}\n`) : console.log(`### ${test}\n`)
            tr.description && console.log(`${testResults[test].description}\n`)
            // get the baseline metrics
            let bl = testResults[test]["baseline"]
            // iterate over the latency buckets and convert them into either formatted text (inc. deltas as needed)
            let v = testResults[test]["results"].map(v => {
                Object.keys(v).map(w => {
                    if (w === 'total' || w === 'name') {
                        return
                    }
                    if (v.name === 'baseline') {
                        v[w] = formatNumber(v[w], false)
                    } else {
                        v[w] = `${formatNumber(v[w], false)}<br>(${formatNumber((v[w] - bl[w]).toFixed(2))})`
                    }
                })
                return [v.name, v.min, v.mean, v['50th'], v['90th'], v['95th'], v['99th'], v.max]
            })
            // finally log the table
            console.log(`${markdownTable([tableHeaders].concat(v))}\n`)
        }

    } catch (error) {
        console.error(error)
        return
    }
}

// get the results from the various test repositories
const convertResults = async (testPath, testName) => {
    try {
        // so within each test folder (e.g. ../tests/static-subgraph/) we fetch all files
        let files = await fs.promises.readdir(path.join(testPath, 'results'))
        for await (let file of files) {
            // then get the results from the 'results' folder within
            let stat = await fs.promises.stat(path.join(path.join(testPath, 'results'), file))
            // check against the regex with the matching group to get the internal name of the test
            if (stat.isFile() && file.match(regex) && file.match(regex).length >= 2) {
                let name = file.match(regex)[1]
                let filepath = path.join(path.join(testPath, 'results'), file)
                // read the JSON export from the taskfile
                let data = await fs.promises.readFile(filepath, 'utf-8')
                let { latencies } = JSON.parse(data)
                // pre-populate the keys for first runs to avoid errors
                if (!testResults[testName]) {
                    testResults[testName] = {}
                }
                if (!testResults[testName]["results"]) {
                    testResults[testName]["results"] = []
                }

                // convert the latencies into millisecond vs nanosecond
                for (let l in latencies) {
                    latencies[l] = (latencies[l] / 1000000).toFixed(2)
                }
                // then push into the results
                testResults[testName]["results"].push({ ...latencies, name })
                if (name === 'baseline') {
                    if (!testResults[testName]['baseline']) {
                        testResults[testName]['baseline'] = {}
                    }
                    testResults[testName]['baseline'] = { ...latencies, name }
                }
            }
        }

        // lastly, check for a description.json file within the root of the test path as it is used to provide "prettier" annotations for the tables
        let descriptionStat = await fs.promises.stat(path.join(testPath, 'description.json'))
        if (descriptionStat.isFile()) {
            let json = await fs.promises.readFile(path.join(testPath, 'description.json'))
            let { title, description } = JSON.parse(json)
            testResults[testName].title = title
            testResults[testName].description = description
        }
    } catch (error) {
        console.error(error)
        return
    }
}

const formatNumber = (number, includePlus = true) => {
    if (!includePlus) {
        return number
    }
    return (number <= 0 ? "" : "+") + number
}
(async () => {
    await getTestResults()
})()

