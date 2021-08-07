<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <a-row>
        <a-col :sm="3" :xs="20">
          <info title="本年支出" :value="analyse.TotalYear" :bordered="true"/>
        </a-col>
        <a-col :sm="3" :xs="20">
          <info title="上月支出" :value="analyse.lastMonth" :bordered="true"/>
        </a-col>
        <a-col :sm="3" :xs="20">
          <info title="本月支出" :value="analyse.currentMonth" :bordered="true" />
        </a-col>
        <a-col :sm="3" :xs="240">
          <info title="本周支出" :value="analyse.currentWeek" :bordered="true" />
        </a-col>
        <a-col :sm="3" :xs="20">
          <info title="上周支出" :value="analyse.lastWeek" :bordered="true"/>
        </a-col>
        <a-col :sm="3" :xs="20">
          <info title="本年日均" :value="analyse.avgYear + ' / ' + analyse.avgYearNoLoad" :bordered="true"/>
        </a-col>
        <a-col :sm="3" :xs="20">
          <info title="本周日均" :value="analyse.avgWeek + ' / ' + analyse.avgWeekNoLoad" :bordered="true"/>
        </a-col>
      </a-row>
    </a-card>
    <a-card style="margin-top: 24px" :bordered="false" title="支出列表">
      <div slot="extra">
        <a-radio-group v-model="status">
          <a-radio-button value="all">全部</a-radio-button>
          <a-radio-button value="processing">进行中</a-radio-button>
          <a-radio-button value="waiting">等待中</a-radio-button>
        </a-radio-group>
      </div>
      <div class="table-page-search-wrapper">
        <a-form layout="inline">
          <a-row :gutter="48">
            <a-col :md="3" :sm="10">
              <a-form-item label="年">
                <a-select v-model="queryParam.year" placeholder="请选择" default-value="0">
                  <a-select-option v-for="item in years" :key="item.value" >
                    {{ item.label }}
                  </a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="3" :sm="10">
              <a-form-item label="月">
                <a-select v-model="queryParam.month" placeholder="请选择" default-value="0">
                  <a-select-option v-for="item in months" :key="item.value" default-value="0">
                    {{ item.label }}
                  </a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="5" :sm="20">
              <a-form-item label="日期" validateStatus="success">
                <a-range-picker v-model="queryParam.date_range" :format="dateFormat" :valueFormat="dateFormat">
                  <template slot="dateRender" slot-scope="current">
                    <div class="ant-calendar-date" :style="getCurrentStyle(current)">
                      {{ current.date() }}
                    </div>
                  </template>
                </a-range-picker>
              </a-form-item>
            </a-col>
            <a-col :md="3" :sm="14">
              <a-form-item label="内容">
                <a-input v-model="queryParam.content" placeholder=""/>
              </a-form-item>
            </a-col>
            <a-col :md="3" :sm="10">
              <a-form-item label="类别">
                <a-select v-model="queryParam.category_id" placeholder="请选择" default-value="0">
                  <a-select-option v-for="item in categories" :key="item.id" >
                    {{ item.name }}
                  </a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <template v-if="advanced">
            </template>
            <a-col :md="!advanced && 8 || 24" :sm="24">
              <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                <a-button type="primary" @click="$refs.table.refresh(true)">查询</a-button>
                <a-button style="margin-left: 8px" @click="() => this.queryParam = {}">重置</a-button>
                <a @click="toggleAdvanced" style="margin-left: 8px">
                  {{ advanced ? '收起' : '展开' }}
                  <a-icon :type="advanced ? 'up' : 'down'"/>
                </a>
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>

      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="handleAdd">新建</a-button>
        <a-dropdown v-action:edit v-if="selectedRowKeys.length > 0">
          <a-menu slot="overlay">
            <a-menu-item key="1"><a-icon type="delete" />删除</a-menu-item>
            <!-- lock | unlock -->
            <a-menu-item key="2"><a-icon type="lock" />锁定</a-menu-item>
          </a-menu>
          <a-button style="margin-left: 8px">
            批量操作 <a-icon type="down" />
          </a-button>
        </a-dropdown>
      </div>

      <s-table
        ref="table"
        size="default"
        rowKey="key"
        :columns="columns"
        :data="loadData"
        :alert="true"
        :rowSelection="rowSelection"
        showPagination="auto"
      >
        <span slot="serial" slot-scope="text, record, index">
          {{ index + 1 }}
        </span>
        <span slot="status" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>
        <span slot="description" slot-scope="text">
          <ellipsis :length="4" tooltip>{{ text }}</ellipsis>
        </span>

        <span slot="action" slot-scope="text, record">
          <template>
            <a @click="handleEdit(record)">修改</a>
            <a-divider type="vertical" />
            <a @click="handleSub(record)">订阅报警</a>
          </template>
        </span>
      </s-table>

      <create-form
        ref="createModal"
        :visible="visible"
        :loading="confirmLoading"
        :model="mdl"
        @cancel="handleCancel"
        @ok="handleOk"
      />
      <step-by-step-modal ref="modal" @ok="handleOk"/>
    </a-card>

  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { STable, Ellipsis } from '@/components'
import { getCostCategories, getCostList } from '@/api/wealth'
import { getAnalyse } from '@/api/analyse'
import Info from './components/Info'

import StepByStepModal from './modules/StepByStepModal'
import CreateForm from './modules/CreateForm'
import { mapActions } from 'vuex'

const columns = [
  {
    title: '#',
    scopedSlots: { customRender: 'serial' }
  },
  {
    title: '日期',
    sorter: true,
    dataIndex: 'occur_date'
  },
  {
    title: '类别',
    sorter: true,
    dataIndex: 'category_name'
  },
  {
    title: '内容',
    dataIndex: 'content'
  },
  {
    title: '金额',
    dataIndex: 'amount',
    needTotal: true,
    sorter: true,
    customRender: (text) => text + ' 元'
  },
  {
    title: '描述',
    dataIndex: 'description'
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '150px',
    scopedSlots: { customRender: 'action' }
  }
]

const statusMap = {
  0: {
    status: 'default',
    text: '关闭'
  },
  1: {
    status: 'processing',
    text: '运行中'
  },
  2: {
    status: 'success',
    text: '已上线'
  },
  3: {
    status: 'error',
    text: '异常'
  }
}

export default {
  name: 'TableList',
  components: {
    Info,
    STable,
    Ellipsis,
    CreateForm,
    StepByStepModal
  },
  data () {
    this.columns = columns
    return {
      yearsModel: null,
      years: [],
      monthsModel: null,
      months: [],
      dateFormat: 'YYYY-MM-DD',
      // create model
      visible: false,
      confirmLoading: false,
      mdl: null,
      // 高级搜索 展开/关闭
      advanced: false,
      // 查询参数
      queryParam: {},
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
        parameter.user_id = 1
        const requestParameters = Object.assign({}, parameter, this.queryParam)
        console.log('loadData request parameters:', requestParameters)
        if (requestParameters.occur_date != null) {
          requestParameters.occur_date = moment(requestParameters.occur_date).format('YYYY-MM-DD')
        }

        return getCostList(requestParameters)
          .then(res => {
            return res.result
          })
      },
      analyse: {
        TotalYear: 0,
        currentMonth: 0,
        currentWeek: 0,
        lastMonth: 0,
        lastWeek: 0,
        avgYear: 0,
        avgYearNoLoad: 0,
        avgWeek: 0,
        avgWeekNoLoad: 0
      },
      categories: [],
      selectedRowKeys: [],
      selectedRows: []
    }
  },
  filters: {
    statusFilter (type) {
      return statusMap[type].text
    },
    statusTypeFilter (type) {
      return statusMap[type].status
    }
  },
  created () {
    this.init()
    this.getCategories()
    this.getAnalyse()
  },
  computed: {
    rowSelection () {
      return {
        selectedRowKeys: this.selectedRowKeys,
        onChange: this.onSelectChange
      }
    }
  },
  methods: {
    moment,
    ...mapActions(['CreateCost', 'UpdateCost']),
    handleAdd () {
      this.mdl = null
      this.visible = true
    },
    handleEdit (record) {
      this.visible = true
      this.mdl = { ...record }
    },
    handleOk () {
      const form = this.$refs.createModal.form
      const { UpdateCost, CreateCost } = this
      this.confirmLoading = true
      form.validateFields((errors, values) => {
        if (!errors) {
          values.User_id = 1
          values.occur_date = moment(values.occur_date).format('YYYY-MM-DD')
          console.log('values', values)
          if (values.id > 0) {
            // 修改 e.g.
            UpdateCost(values).then(res => {
              this.visible = false
              this.confirmLoading = false
              // 重置表单数据
              form.resetFields()
              // 刷新表格
              this.$refs.table.refresh()

              this.$message.info('修改成功')
            })
          } else {
            // 新增
           CreateCost(values).then(res => {
              this.visible = false
              this.confirmLoading = false
              // 重置表单数据
              form.resetFields()
              // 刷新表格
              this.$refs.table.refresh()

              this.$message.info('新增成功')
            })
          }
        } else {
          this.confirmLoading = false
        }
      })
    },
    handleCancel () {
      this.visible = false

      const form = this.$refs.createModal.form
      form.resetFields() // 清理表单数据（可不做）
    },
    handleSub (record) {
      if (record.status !== 0) {
        this.$message.info(`${record.no} 订阅成功`)
      } else {
        this.$message.error(`${record.no} 订阅失败，规则已关闭`)
      }
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    getCategories () {
      var parameter = {}
      parameter.user_id = 1
      getCostCategories(parameter).then(res => {
        if (res.result.data.length > 0) {
          this.categories = res.result.data
        }
      }).catch((err) => {
        console.log('category list', err)
      })
    },
    getAnalyse () {
      var parameter = {}
      parameter.user_id = 1
      getAnalyse(parameter).then(res => {
        this.analyse.currentMonth = res.result.current_month
        this.analyse.currentWeek = res.result.current_week
        this.analyse.lastMonth = res.result.last_month
        this.analyse.lastWeek = res.result.last_week
        this.analyse.avgYear = res.result.avg_current_year
        this.analyse.avgYearNoLoad = res.result.avg_current_year_no_load
        this.analyse.avgWeek = res.result.avg_current_week
        this.analyse.avgWeekNoLoad = res.result.avg_current_week_no_load
        this.analyse.TotalYear = res.result.total_year
      }).catch((err) => {
        console.log('category list', err)
      })
    },
    getCurrentStyle (current, today) {
      const style = {}
      if (current.date() === 1) {
        style.border = '1px solid #1890ff'
        style.borderRadius = '50%'
      }
      return style
    },
    init () {
      var myDate = new Date()
      var year = myDate.getFullYear() // 获取当前年
      // var month = myDate.getMonth() + 1 // 获取当前月

      this.initSelectYear(year)
      this.initSelectMonth()
      // this.queryParam.year = year
      // this.queryParam.month = month
    },
    initSelectYear (year) {
      this.years = []
      this.years.push({ value: 0, label: '全部' })
      for (let i = 0; i < 30; i++) {
        this.years.push({ value: (year - i), label: (year - i) + '年' })
      }
    },
    initSelectMonth () {
      this.months = []
      this.months.push({ value: 0, label: '全部' })
        for (let i = 1; i <= 12; i++) {
          this.months.push({ value: i, label: i + '月' })
        }
    },
    resetSearchForm () {
      this.queryParam = {
        date: moment(new Date())
      }
    }
  }
}
</script>
