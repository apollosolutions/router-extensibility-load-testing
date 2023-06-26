import { markdownTable } from 'markdown-table'
import fs from 'fs'
import path from 'path'

const RESULTS_JSON_REGEX = /results_(?<name>.*)\.json$/

// get results from results_*.json files for each test type
const buildTestResults = async () => {
  const testsPath = path.join('..', 'tests')
  const testResults = {}

  const testDirectories = await fs.promises.readdir(testsPath)
  for (let type of testDirectories) {
    const testTypePath = path.join(testsPath, type)
    testResults[type] = await parseResultsForTestType(testTypePath, type)
  }

  return testResults
}

// parse the results from the various test type directories
const parseResultsForTestType = async (testTypePath, altTitle) => {
  const testType = { baseline: {}, results: [], title: altTitle }

  // check for a meta.json file within the root of the test path to get "prettier" table annotations
  const metaPath = path.join(testTypePath, 'meta.json')
  try {
    const data = await fs.promises.readFile(metaPath)
    const { description, title } = JSON.parse(data)
    testType.title = title
    testType.description = description
  } catch (err) {
    console.error(`Metadata for ${altTitle} not found`)
  }

  // fetch all files within each test's results dir (e.g. ../tests/static-subgraph/results/)
  const resultsPath = path.join(testTypePath, 'results')
  try {
    const files = await fs.promises.readdir(resultsPath)
    for (let file of files) {
      const filePath = path.join(resultsPath, file)
      const { isJson, ...result } = await parseJsonTestResult(filePath)

      // append the results from json files
      if (isJson) {
        testType.results.push(result)
        if (result.name === 'baseline') {
          testType.baseline = result
        }
      }
    }
  } catch (err) {
    console.error(`Results for ${altTitle} not found`)
  }

  return testType
}

// get the results from the results_*.json files
const parseJsonTestResult = async (file) => {
  let result = { isJson: false }

  // check against the regex with the matching group to get the internal name of the test run
  const match = file.match(RESULTS_JSON_REGEX)
  if (!!match?.groups?.name) {
    try {
      // read the JSON export from the test task
      const data = await fs.promises.readFile(file, 'utf-8')
      const { latencies, success } = JSON.parse(data)

      result = { ...latencies, isJson: true, name: match.groups.name, success }
    } catch (err) {
      console.error(`Failed to read results file: ${file}`)
    }
  }

  return result
}

// convert results into markdown and log to console
const printResults = (testResults) => {
  Object.values(testResults).forEach(test => {
    console.log(`### ${test.title}\n`)
    test.description && console.log(`${test.description}\n`)

    // iterate over the latency buckets and convert them into formatted text with deltas
    const tableRows = test.results.map(result => {
      const isBaseline = result.name === 'baseline'

      return [
        result.name,
        formatLatency(result.min, test.baseline.min, isBaseline),
        formatLatency(result.mean, test.baseline.mean, isBaseline),
        formatLatency(result['50th'], test.baseline['50th'], isBaseline),
        formatLatency(result['90th'], test.baseline['90th'], isBaseline),
        formatLatency(result['95th'], test.baseline['95th'], isBaseline),
        formatLatency(result['99th'], test.baseline['99th'], isBaseline),
        formatPercentage(result.success),
      ]
    })

    // print the table
    const table = markdownTable([
      ['Type', 'Min (ms)', 'Mean (ms)', 'p50 (ms)', 'p90 (ms)', 'p95 (ms)', 'p99 (ms)', 'Success Rate'],
      ...tableRows,
    ], { align: ['left', 'center', 'center', 'center', 'center', 'center', 'center', 'center'] })
    console.log(`${table}\n`)
  })
}

// format latency values with delta from baseline
const formatLatency = (value, baseline, isBaseline) => {
  // convert latency values from nanoseconds to milliseconds
  let formatted = (value / 1000000).toFixed(2)

  if (!isBaseline) {
    const delta = (value - baseline) / 1000000
    formatted = `${formatted}<br>${delta > 0 ? '+' : ''}${delta.toFixed(2)}`
  }

  return formatted
}

const formatPercentage = (value) => {
  return `${(value * 100).toFixed(0)}%`
}

(async () => {
  try {
    const testResults = await buildTestResults()
    printResults(testResults)
  } catch (err) {
    console.error(err)
    process.exitCode = 1
  }
})()
