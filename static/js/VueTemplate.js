//Vue组件
$(function () {
    Vue.component('vadd-text', {
        props: ['txt_index', 'txt_item'],
        template: `
                        <div class="input-group input-icon right">
                            <span class= "input-group-addon">
                                <i :class="txt_item.Icon"></i>
                            </span>
                            <input type="text" v-model="txt_item.FieldValue" :placeholder="txt_item.Placeholder" class="form-control1 icon"/>
                        </div>
                        `
    });
    Vue.component('vadd-checkbox', {
        props: ['ck_index', 'ck_item'],
        template: `
                         <div>
                             <div class="checkbox-inline"  v-for="option in ck_item.FieldSelect"><label><input type="radio" v-model="ck_item.FieldValue" :value="option.value">{{option.text}}</label></div>
                         </div>
                        `
    });
    Vue.component('vadd-textarea', {
        props: ['txtarea_index', 'txtarea_item'],
        template: `
                         <div>
                             <textarea class="form-control1" cols="50" rows="4" style="height:70px;" v-model="txtarea_item.FieldValue"></textarea>
                         </div>
                          `
    });
    Vue.component('vadd-select', {
        props: ['select_index', 'select_item'],
        template: `
                         <select class="form-control1" v-model="select_item.FieldValue">
                             <option v-for="option in select_item.FieldSelect" :value="option.value" >{{option.text}}</option>
                         </select>
                          `
    });
})
