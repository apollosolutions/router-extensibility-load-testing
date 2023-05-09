import { markdownTable } from "markdown-table";
import fs from 'fs'
import path from 'path'

const regex = /results_(.*).json/
const tableHeaders = ['Type', 'Min', 'Mean', 'p50', 'p90', 'p95', 'p99', 'Max']
let testResults = {}

// get the test results from the directory
const getTestResults = async () => {
    try {
        let testFolder = '../tests'
        let files = await fs.promises.readdir(testFolder)
        for await (let file of files) {
            let stat = await fs.promises.stat(path.join(testFolder, file))
            if (stat.isDirectory()) {
                await convertResults(path.join(testFolder, file), file)
            }
        }

        for (let test in testResults) {
            console.log(`### ${test}\n`)
            console.log(markdownTable([tableHeaders].concat(testResults[test])) + '\n')
        }

    } catch (error) {
        console.error(error)
        return
    }
}

// get the results from the various test repositories
const convertResults = async (testPath, testName) => {
    try {
        let files = await fs.promises.readdir(path.join(testPath, 'results'))
        for await (let file of files) {
            let stat = await fs.promises.stat(path.join(path.join(testPath, 'results'), file))
            if (stat.isFile() && file.match(regex) && file.match(regex).length >= 2) {
                let name = file.match(regex)[1]
                let filepath = path.join(path.join(testPath, 'results'), file)
                let data = await fs.promises.readFile(filepath, 'utf-8')
                let { latencies } = JSON.parse(data)
                if (!testResults[testName]) {
                    testResults[testName] = []
                }
                for (let l in latencies) {
                    latencies[l] = (latencies[l] / 1000000).toFixed(2) + 'ms' // convert latencies to ms vs ns
                }
                testResults[testName].push([name, latencies.min, latencies.mean, latencies['50th'], latencies['90th'], latencies['95th'], latencies['99th'], latencies.max])
            }
        }

    } catch (error) {
        console.error(error)
        return
    }
}

(async () => {
    await getTestResults()
})()
