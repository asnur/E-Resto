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
