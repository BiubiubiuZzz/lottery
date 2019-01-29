

$(function() {
	
	// 奖类管理
	var $obj=$(".award_manager table tr th,.award_manager table tr td");
	var $optObj=$(".award_manager table .aClass,.award_manager table .aType,.award_manager table .aSmallclass");
	$optObj.click(function() {
		$obj.removeClass('selected');
		$(this).addClass('selected');
	});

	
	// 当前selected对象弹窗
	// 新增
	$('.tools .add').click(function() {
		if ($('.selected').hasClass('aClass')) {
			$('.mengban,.tbox-addclass').fadeIn();
		} 
		else if($('.selected').hasClass('aType')) {
			$('.mengban,.tbox-addtype').fadeIn();
		} 
		else if($('.selected').hasClass('aSmallclass')) {
			$('.mengban,.tbox-addsmallclass').fadeIn();
		} 
		else{
			alert('请先点击选择相应级别再点击新增');
		};	
	});

	// 修改
	// 读取当前selected数据填入input，修改提交update数据库by字段id

	$('.tools .modify').click(function() {
		if ($('.selected').hasClass('aClass')) {
			alert('不支持修改类别，表问为什么，十万个为什么');
		} 
		else if($('.selected').hasClass('aType')) {
			$('.mengban,.tbox-modifytype').fadeIn();
		} 
		else if($('.selected').hasClass('aSmallclass')) {
			$('.mengban,.tbox-modifysmallclass').fadeIn();
		} 
		else{
			alert('请先点击选择类型或小类再点击修改');
		};
	});




	// 删除
	$('.tools .delete').click(function() { 
	  
	    function del(){
	  	var r=confirm("请仔细确认，您要删除的是:"+$('.selected').val() +"，确定要删除么？")
	  	if (r==true){
	    	$('.selected').parent('tr').remove(); 
	  	}
	  	else{};
	  }

	    if     ($('.selected').hasClass('aClass')) {
	  		// 所选的类别/类型下有奖品的话，不能直接删除
	  		// if () {alert("所选的类别/类型下有奖品的话，不能直接删除");} else{del();};
		} 
		else if($('.selected').hasClass('aType')) {
			del();
		} 
		else if($('.selected').hasClass('aSmallclass')) {
			del();
		} 
		else{
			alert('请先点击选择再点删除');
		};
	});

	$('.queryclass').click(function() {
	    $('.award_manager .p_txt').show();
	});

	//自定义二次确认弹窗
	// $('.tools .delete').click(function(){
	// 	$('.confirmbox').fadeIn();
	// })
	// $('.confirmdel,.btncancel').click(function(){
	// 	$('.confirmbox').fadeOut();
	// })

	// 关闭弹窗
	$('.btncancel,.btnsave').bind('click', function() {
		$('.mengban,.editwrap,.confirmbox').fadeOut();
	});

	// 活动管理
		// 新增奖品
	$('.aw_newadd').click(function() {
		$('.mengban,.addaw_box').fadeIn();
	});
		// 追加数量
	$('td .lateadd').click(function() {
		$('.mengban,.addaw_box1').fadeIn();
	});

	$('.btncancel').click(function() {
		$('.mengban,.addaw_box,.addaw_box1').fadeOut();
	});


	// 奖品管理表格字段滑过显示完整

	$('.jpsqtable .td8').mouseover(function(event) {
		$(this).attr('title', $(this).html());
	});

	// 奖池管理页
	// 前台样式

	$('.aw_edit').click(function(event) {
		/* Act on the event */
	});

	$('.aw_del').click(function(event) {
		/* Act on the event */
	});

	$('.aw_state').click(function(){
		if ($(this).html()=='启用') {
			$(this).html('禁用');
		} 
		else{
			$(this).html('启用');
		};
	});

	$('.aw_log').click(function(event) {
		/* Act on the event */
	});

	$('.aw_couponlog').click(function(event) {
		/* Act on the event */
	});

	var i=0;
	// 输入框前台check
	function check(){
	    var $v = document.getElementById('target');
	    if(v == ''){
	        alert('值不能为空');
	        return false;
	    }
	    if(v.length > 10){
	        alert('值长度不能超过10');
	        return false;
	    }
	    re = new RegExp("^[a-zA-Z\u4e00-\u9fa5]+$"); 
	    if (!re.test(v)) {
	        alert('只能输入汉字和字母');
	        return false;
	    }
	    return true;
	    console.log('输入已验证');
	}
	$('.queryclass').bind('click', function(event) {
		i++;if (i==15) {console.log('\u597d\u5427\u4f60\u8fd9\u4e48\u65e0\u804a\u5c45\u7136\u70b9\u4e86\u0031\u0035\u6b21\u003d\u005f\u003d')}
		else if(i==20){console.log('\u5c40\u65b9\u662f\u9017\u903c\u003d\u005f\u003d');};
	});
})
	// jquery.validate.js
		//$('.js-form') 待验证表单
	    //submitHandler 验证通过后执行的函数
		// $('.js-form').validate({ 		
		// 	submitHandler: function (form) {
	 //            var $form = $(form),
	 //                data = $form.serialize();    //序列化表单数据
	 //            //验证通过后如ajax post      
	 //            $.post('js/test.json',{data:data},function(d){
	 //                if(d.Flag){}
	 //                	else{}
	 //            },'json');
		// 	}
		// });
		
		
//		超出字段...显示,自动写入title值
// function txtAllShow(){
// 	var obj=$('.txtover')
// 	$(obj).mouseover(function() {
//        $(this).attr('title', $(this).html());
//     });
// }
// txtAllShow();

	