import React from "react";

const Customers = ({customers}) =>{




    return(

        <table className="table">
        <thead>
          <tr >
            <th scope="col">Cliente ID</th>
            <th scope="col">Compañía</th>
            <th scope="col">Nombre</th>
            <th scope="col">Cargo</th>
            <th scope="col">Dirección</th>
            <th scope="col">Ciudad</th>
            <th scope="col">Región</th>
            <th scope="col">Código Postal</th>
            <th scope="col">País</th>
            <th scope="col">Teléfono</th>
          </tr>
        </thead>
        <tbody>
            {customers.map((customer)=>(
              
              <tr key={customer.customer_id}>
              <th scope="row">{customer.customer_id}</th>
              <td>{customer.company_name}</td>
              <td>{customer.contact_name}</td>
              <td>{customer.contact_title}</td>
              <td>{customer.address}</td>
              <td>{customer.city}</td>
              <td>{customer.region}</td>
              <td>{customer.postal_code}</td>
              <td>{customer.country}</td>
              <td>{customer.phone}</td>
            </tr>
           
            ))}
        </tbody>
  </table>
    )
};

export default Customers

  


  
