<template>
  <a-modal
    title="赚钱啦"
    :width="640"
    :visible="visible"
    :confirmLoading="loading"
    @ok="() => { $emit('ok') }"
    @cancel="() => { $emit('cancel') }"
  >
    <a-spin :spinning="loading">
      <a-form :form="form" v-bind="formLayout">
        <!-- 检查是否有 id 并且大于0，大于0是修改。其他是新增，新增不显示主键ID -->
        <a-form-item v-show="model && model.id > 0" label="主键ID">
          <a-input v-decorator="['id', { initialValue: 0 }]" disabled />
        </a-form-item>

        <a-form-item label="类型" validateStatus="success" >
          <a-select :default-value="{ key: 0 }" v-decorator="['type', {rules: [{ required: true, message: '请选择类型' }]}]">
            <a-select-option v-for="item in types" :key="item.id" >
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item label="内容" validateStatus="warning" >
          <a-input placeholder="内容" v-decorator="[ 'content', {rules: [{ required: true, message: '请输入内容' }]} ]"></a-input>
        </a-form-item>

        <a-form-item label="金额" help="请填写支出金额（单位：元）">
          <a-input placeholder="金额" v-decorator="[ 'amount', {rules: [{ required: true, message: '请输入金额' }]} ]"></a-input>
        </a-form-item>

        <a-form-item label="描述">
          <a-input v-decorator="['description']" />
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import pick from 'lodash.pick'
import moment from 'moment'

// 表单字段
const fields = ['type', 'content', 'amount', 'description', 'id']

const types = [
  {
    id: 1,
    name: '活期存款'
  },
  {
    id: 2,
    name: '定期存款'
  },
  {
    id: 3,
    name: '余额宝'
  },
  {
    id: 4,
    name: '支付宝'
  },
  {
    id: 5,
    name: '微信'
  },
  {
    id: 6,
    name: '零钱通'
  },
  {
    id: 7,
    name: '现金'
  }
]
export default {
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    loading: {
      type: Boolean,
      default: () => false
    },
    model: {
      type: Object,
      default: () => null
    }
  },
  data () {
    this.formLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      }
    }
    return {
      types: types,
      form: this.$form.createForm(this)
    }
  },
  created () {
    console.log('custom modal created')

    // 防止表单未注册
    fields.forEach(v => this.form.getFieldDecorator(v))

    // 当 model 发生改变时，为表单设置值
    this.$watch('model', () => {
      this.model && this.form.setFieldsValue(pick(this.model, fields))
    })
  },
  methods: {
    moment
  }
}
</script>
