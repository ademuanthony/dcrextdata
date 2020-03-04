import { Controller } from 'stimulus'
import axios from 'axios'
import {
  legendFormatter,
  barChartPlotter,
  hide,
  show,
  setActiveOptionBtn,
  options,
  showLoading,
  hideLoading,
  selectedOption, insertOrUpdateQueryParam, updateQueryParam, updateZoomSelector
} from '../utils'
import TurboQuery from '../helpers/turbolinks_helper'
import Zoom from '../helpers/zoom_helper'
import { animationFrame } from '../helpers/animation_helper'

const Dygraph = require('../../../dist/js/dygraphs.min.js')

export default class extends Controller {
  static get targets () {
    return [
      'nextPageButton', 'previousPageButton', 'tableBody', 'rowTemplate',
      'totalPageCount', 'currentPage', 'btnWrapper', 'tableWrapper', 'chartsView',
      'chartWrapper', 'viewOption', 'labels', 'viewOptionControl', 'messageView',
      'chartDataTypeSelector', 'chartDataType', 'chartOptions', 'labels', 'selectedMempoolOpt',
      'selectedNumberOfRows', 'numPageWrapper', 'loadingData',
      'zoomSelector', 'zoomOption'
    ]
  }

  initialize () {
    this.currentPage = parseInt(this.currentPageTarget.getAttribute('data-current-page'))
    if (this.currentPage < 1) {
      this.currentPage = 1
    }

    this.query = new TurboQuery()
    this.settings = TurboQuery.nullTemplate(['chart', 'zoom', 'scale', 'bin', 'axis', 'dataType', 'page', 'view-option'])
    this.settings.chart = this.settings.chart || 'mempool'

    this.zoomCallback = this._zoomCallback.bind(this)
    this.drawCallback = this._drawCallback.bind(this)

    this.dataType = this.chartDataTypeTarget.getAttribute('data-initial-value')

    this.selectedViewOption = this.viewOptionControlTarget.getAttribute('data-initial-value')
    if (this.selectedViewOption === 'chart') {
      this.setChart()
    } else {
      this.setTable()
    }
  }

  setTable () {
    this.selectedViewOption = 'table'
    setActiveOptionBtn(this.selectedViewOption, this.viewOptionTargets)
    hide(this.chartWrapperTarget)
    hide(this.messageViewTarget)
    hide(this.chartDataTypeSelectorTarget)
    hide(this.zoomSelectorTarget)
    show(this.tableWrapperTarget)
    show(this.numPageWrapperTarget)
    show(this.btnWrapperTarget)
    this.nextPage = this.currentPage
    this.fetchData(this.selectedViewOption)
    insertOrUpdateQueryParam('view-option', this.selectedViewOption)
  }

  setChart () {
    this.selectedViewOption = 'chart'
    hide(this.btnWrapperTarget)
    hide(this.tableWrapperTarget)
    hide(this.messageViewTarget)
    setActiveOptionBtn(this.selectedViewOption, this.viewOptionTargets)
    setActiveOptionBtn(this.dataType, this.chartDataTypeTargets)
    show(this.chartDataTypeSelectorTarget)
    show(this.zoomSelectorTarget)
    hide(this.numPageWrapperTarget)
    show(this.chartWrapperTarget)
    this.fetchData(this.selectedViewOption)
    updateQueryParam('view-option', this.selectedViewOption)
  }

  setDataType (event) {
    this.dataType = event.currentTarget.getAttribute('data-option')
    setActiveOptionBtn(this.dataType, this.chartDataTypeTargets)
    this.fetchData('chart')
    insertOrUpdateQueryParam('chart-data-type', this.dataType)
  }

  numberOfRowsChanged () {
    this.selectedNumberOfRowsberOfRows = this.selectedNumberOfRowsTarget.value
    this.fetchData(this.selectedViewOption)
    insertOrUpdateQueryParam('records-per-page', this.selectedNumberOfRowsberOfRows)
  }

  loadPreviousPage () {
    this.nextPage = this.currentPage - 1
    this.fetchData(this.selectedViewOption)
    insertOrUpdateQueryParam('page', this.nextPage)
  }

  loadNextPage () {
    this.nextPage = this.currentPage + 1
    this.fetchData(this.selectedViewOption)
    insertOrUpdateQueryParam('page', this.nextPage)
  }

  fetchData (display) {
    let url
    let elementsToToggle = [this.tableWrapperTarget, this.chartWrapperTarget]
    showLoading(this.loadingDataTarget, elementsToToggle)

    if (display === 'table') {
      this.selectedNumberOfRowsberOfRows = this.selectedNumberOfRowsTarget.value
      url = `/getmempool?page=${this.nextPage}&records-per-page=${this.selectedNumberOfRowsberOfRows}&view-option=${this.selectedViewOption}`
    } else {
      url = `/mempoolcharts?chart-data-type=${this.dataType}&view-option=${this.selectedViewOption}`
    }

    const _this = this
    axios.get(url).then(function (response) {
      let result = response.data
      if (display === 'table' && result.message) {
        hideLoading(_this.loadingDataTarget, [_this.tableWrapperTarget])
        let messageHTML = ''
        messageHTML += `<div class="alert alert-primary">
                       <strong>${result.message}</strong>
                  </div>`

        _this.messageViewTarget.innerHTML = messageHTML
        show(_this.messageViewTarget)
        hide(_this.tableBodyTarget)
        hide(_this.btnWrapperTarget)
      } else if (display === 'table' && result.mempoolData) {
        hideLoading(_this.loadingDataTarget, [_this.tableWrapperTarget])
        hide(_this.messageViewTarget)
        show(_this.tableBodyTarget)
        show(_this.btnWrapperTarget)
        _this.totalPageCountTarget.textContent = result.totalPages
        _this.currentPageTarget.textContent = result.currentPage

        _this.currentPage = result.currentPage
        if (_this.currentPage <= 1) {
          _this.currentPage = result.currentPage
          hide(_this.previousPageButtonTarget)
        } else {
          show(_this.previousPageButtonTarget)
        }

        if (_this.currentPage >= result.totalPages) {
          hide(_this.nextPageButtonTarget)
        } else {
          show(_this.nextPageButtonTarget)
        }

        _this.displayMempool(result.mempoolData)
      } else {
        hideLoading(_this.loadingDataTarget, [_this.chartWrapperTarget])
        _this.plotGraph(result)
      }
    }).catch(function (e) {
      hideLoading(_this.loadingDataTarget)
      console.log(e) // todo: handle error
    })
  }

  displayMempool (data) {
    const _this = this
    this.tableBodyTarget.innerHTML = ''

    data.forEach(item => {
      const exRow = document.importNode(_this.rowTemplateTarget.content, true)
      const fields = exRow.querySelectorAll('td')

      fields[0].innerText = item.time
      fields[1].innerText = item.number_of_transactions
      fields[2].innerText = item.size
      fields[3].innerHTML = item.total_fee.toFixed(8)

      _this.tableBodyTarget.appendChild(exRow)
    })
  }

  selectedZoom () { return selectedOption(this.zoomOptionTargets) }

  setZoom (e) {
    var target = e.srcElement || e.target
    var option
    if (!target) {
      let ex = this.chartsView.xAxisExtremes()
      option = Zoom.mapKey(e, ex, 1)
    } else {
      option = target.dataset.option
    }
    setActiveOptionBtn(option, this.zoomOptionTargets)
    if (!target) return // Exit if running for the first time
    this.validateZoom()
  }

  async validateZoom () {
    await animationFrame()
    await animationFrame()
    let oldLimits = this.limits || this.chartsView.xAxisExtremes()
    this.limits = this.chartsView.xAxisExtremes()
    var selected = this.selectedZoom()
    if (selected) {
      this.lastZoom = Zoom.validate(selected, this.limits, 1, 1)
    } else {
      this.lastZoom = Zoom.project(this.settings.zoom, oldLimits, this.limits)
    }
    if (this.lastZoom) {
      this.chartsView.updateOptions({
        dateWindow: [this.lastZoom.start, this.lastZoom.end]
      })
    }
    if (selected !== this.settings.zoom) {
      this._zoomCallback(this.lastZoom.start, this.lastZoom.end)
    }
    await animationFrame()
    this.chartsView.updateOptions({
      zoomCallback: this.zoomCallback,
      drawCallback: this.drawCallback
    })
  }

  _zoomCallback (start, end) {
    this.lastZoom = Zoom.object(start, end)
    this.settings.zoom = Zoom.encode(this.lastZoom)
    let ex = this.chartsView.xAxisExtremes()
    let option = Zoom.mapKey(this.settings.zoom, ex, 1)
    setActiveOptionBtn(option, this.zoomOptionTargets)
  }

  _drawCallback (graph, first) {
    if (first) return
    var start, end
    [start, end] = this.chartsView.xAxisRange()
    if (start === end) return
    if (this.lastZoom.start === start) return // only handle slide event.
    this._zoomCallback(start, end)
  }

  // exchange chart
  plotGraph (exs) {
    const _this = this
    if (exs.error) {
      this.drawInitialGraph()
    } else {
      let chartData = exs.mempoolchartData
      let csv = ''
      switch (this.dataType) {
        case 'size':
          this.title = 'Size'
          csv = 'Date,Size\n'
          break
        case 'total_fee':
          this.title = 'Total Fee'
          csv = 'Date,Total Fee\n'
          break
        default:
          this.title = '# of Transactions'
          csv = 'Date,# of Transactions\n'
          break
      }
      let minDate, maxDate

      chartData.forEach(mp => {
        let date = new Date(mp.time)
        if (minDate === undefined || new Date(mp.time) < minDate) {
          minDate = new Date(mp.time)
        }

        if (maxDate === undefined || new Date(mp.time) > maxDate) {
          maxDate = new Date(mp.time)
        }

        let record
        if (_this.dataType === 'size') {
          record = mp.size
        } else if (_this.dataType === 'total_fee') {
          record = mp.total_fee
        } else {
          record = mp.number_of_transactions
        }
        csv += `${date},${record}\n`
      })

      _this.chartsView = new Dygraph(
        _this.chartsViewTarget,
        csv,
        {
          legend: 'always',
          includeZero: true,
          dateWindow: [minDate, maxDate],
          legendFormatter: legendFormatter,
          plotter: barChartPlotter,
          digitsAfterDecimal: 8,
          labelsDiv: _this.labelsTarget,
          ylabel: _this.title,
          xlabel: 'Date',
          labelsUTC: true,
          labelsKMB: true,
          maxNumberWidth: 10,
          showRangeSelector: true,
          axes: {
            x: {
              drawGrid: false
            },
            y: {
              axisLabelWidth: 90
            }
          }
        }
      )

      _this.validateZoom()
      updateZoomSelector(_this.zoomOptionTargets, minDate, maxDate)
    }
  }

  drawInitialGraph () {
    var extra = {
      legendFormatter: legendFormatter,
      labelsDiv: this.labelsTarget,
      ylabel: this.title,
      xlabel: 'Date',
      labelsUTC: true,
      labelsKMB: true,
      axes: {
        x: {
          drawGrid: false
        }
      }
    }

    this.chartsView = new Dygraph(
      this.chartsViewTarget,
      [[0, 0]],
      { ...options, ...extra }
    )
  }
}
