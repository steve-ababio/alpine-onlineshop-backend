package checkoutcontroller

import (
	"fmt"
	"model/checkout"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"

)

  type userItemsInfo struct{
    ProductID string `json:"productID"`;
    Quantity int64 `json:"qty"`; 
  } 
  type productid struct{
    ProductID string `json:"productID`
  }
   var req struct{
    Items []userItemsInfo `json:"iteminfo"`;
   }
  
  func filterdbitems(itemdb []*checkout.Itemdbinfo)map[string]int64{
      itemdbfiltered := make(map[string]int64);
      for _,v := range itemdb{
        itemdbfiltered[v.ProductID] = int64(v.Price);
      }
      return itemdbfiltered;
  }
  func calculatetotalamt(itemdb []*checkout.Itemdbinfo,itemuser []userItemsInfo)int64{
      totalamt := int64(0) ;
      itemdbfiltered := filterdbitems(itemdb);
      for _,item := range itemuser{
        if _,ok := itemdbfiltered[item.ProductID]; ok{
          totalamt += (itemdbfiltered[item.ProductID] * item.Quantity);
        }
    }
    return int64(totalamt) ;
  }
  func createintentparams(totalamt int64) *stripe.PaymentIntentParams{
    params := &stripe.PaymentIntentParams{
      Amount: stripe.Int64(totalamt),
      Currency: stripe.String(string(stripe.CurrencyUSD)),
      PaymentMethodTypes: stripe.StringSlice([]string{
        "card",
      }),
      StatementDescriptor: stripe.String("customer descriptor"),
  	};
    params.AddMetadata("integration_test","accept payment");
    return params;
  }
  func Checkout(context *gin.Context){
      context.Header("Content-Type", "application/json")
      err := context.BindJSON(&req);
      fmt.Println(req)
      product_ids := []string{""};

      for _,i := range req.Items{
        product_ids = append(product_ids,i.ProductID);
      }

     var itemresults []*checkout.Itemdbinfo = checkout.Getproducts(product_ids);

      totalamt := calculatetotalamt(itemresults,req.Items);
      if err != nil{
        http.Error(context.Writer,err.Error(),http.StatusInternalServerError);
      }
      params := createintentparams(totalamt);
      pi,err := paymentintent.New(params)

      fmt.Printf("pi.New: %v", pi.ClientSecret)
      if err != nil {
          http.Error(context.Writer, err.Error(), http.StatusInternalServerError)
          fmt.Printf("pi.New: %v", err)
          return
      }
      context.JSON(http.StatusOK,struct{ClientSecret string `json:"clientsecret"`}{ClientSecret:pi.ClientSecret})
  }

  
