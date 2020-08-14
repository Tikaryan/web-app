$(document).on('change', '#email',function(){
       let loginId = $('#email').val()
        if( loginId != "")
        {
               $.ajax({url: "checkUser",type: "POST",data: {"loginid": loginId},
               dataType: "JSON",
               success: function(res)
                      {
                            if (res == true ){
                                   $('#email').val('')
                                   $('#email').addClass('is-invalid')
                                   // $('#email').after('<p style="color:#fe5461; font-size:12px;">Email already is in Use.</p>')
                             }
                      }
                      
               })
        }
 })

 $(document).on('blur', '#form', function(){
        $('#form input[type=email]').each(function(){
               $(this).removeClass('is-invalid')
        })
 })