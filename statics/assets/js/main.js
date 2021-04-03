/* ----------------- Start Document ----------------- */
(function($) {
    $(document).ready(function() {
        $(document).on('click', "a[name='insta_modal']", function() {
            $pid = $(this).attr('pid');
            $input_element = $(`input[name='photo_content'][pid='${$pid}']`)[0];
            $caption = $input_element.value;
            $img_element = $(`img[name='insta_modal'][pid='${$pid}']`)[0];
            $src = $img_element.src;
            $("#caption_insta").text($caption);
            $("#insta_img")[0].src = $src;
            $("#insta_img2")[0].src = $src;
            $("#download_insta").attr("src", $src)
        });
        $(document).on('click', "a[id='download_insta']", function() {
            $src = $(this).attr('src');
            console.log($src);
            var today = new Date();
            var date = today.getFullYear() + '_' + (today.getMonth() + 1) + '_' + today.getDate();
            var time = today.getHours() + "_" + today.getMinutes() + "_" + today.getSeconds();
            var dateTime = date + '_' + time;
            // 
            var link = document.createElement('a');
            link.href = $src;
            link.download = 'Download.jpg';
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);



        });
    });

})(this.jQuery);