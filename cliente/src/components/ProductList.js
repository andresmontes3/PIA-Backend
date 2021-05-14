import React from "react";

const Products = ({products}) =>{




    return(

        <table className="table">
        <thead>
          <tr >
            <th scope="col">Producto ID</th>
            <th scope="col">Nombre</th>
            <th scope="col">Cantidad por unidad</th>
            <th scope="col">Precio por unidad</th>
            <th scope="col">Unidades en stock</th>
            <th scope="col">Descontinuado</th>
          </tr>
        </thead>
        <tbody>
            {products.map((product)=>(
              
              <tr key={product.product_id}>
              <th scope="row">{product.product_id}</th>
              <td>{product.product_name}</td>
              <td>{product.units_in_stock}</td>
              <td>{product.unit_price}</td>
              <td>{product.units_in_stock}</td>
              <td>{product.discontinued.toString()}</td>
            </tr>
           
            ))}
        </tbody>
  </table>
    )
};

export default Products

  


  
