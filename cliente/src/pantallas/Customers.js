import React from "react";
import CustomerList from '../components/CustomerList';


class Customers extends React.Component{

    state = {
      customers: [],
        customer_id: '',
        company_name: '',
        contact_name: '',
        contact_title: '',
        address: '',
        city: '',
        region: '',
        postal_code: '',
        phone: '',
     };
  
     handleChanges = async(e)=>{
       const target = e.target
        await this.setState({
         [target.name]:target.value
       });
     }
  
  
    
     reqDelete=()=>{
      fetch('http://localhost:1323/customers?customer_id='+this.state.customer_id,{
        method:'delete'
     })
     .then((data) => {
      this.reqGet()
       console.log(data)
     })
     .catch(console.log)
     }
  
     reqGet=()=>{
      fetch('http://localhost:1323/customers',{
        method:'get'
     })
     .then(res => res.json())
     .then((data) => {
       console.log(data)
       this.setState({ customers: data })
     })
     .catch(console.log)
     }
  
     reqPost=async()=>{
      await fetch('http://localhost:1323/customers',{
        method:'post',
        headers: {
          'Accept': 'application/json, text/plain, */*',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(
          {
            customer_id : this.state.customer_id,
            company_name : this.state.company_name,
            contact_name : this.state.contact_name,
            contact_title : this.state.contact_title,
            address : this.state.address,
            city : this.state.city,
            region : this.state.region,
            country : this.state.country,
            postal_code : this.state.postal_code,
            phone : this.state.phone
          })
     })
     .then((data) => {
       this.reqGet()
       console.log(data)
     })
     .catch(console.log)
     }
   
     componentDidMount() {
       this.reqGet();
     }
     
  
     render(){
  
      return (
  
  
        <div className="container">
        <h2>Lista Clientes</h2>
        <CustomerList customers={this.state.customers}/>
        <h3>Información Nuevo Cliente</h3>
        <form>   
              <label htmlFor="newcustomer" className="form-label">Compañía</label>
              <input type="text" className="form-control" name="company_name" id="company_name" onChange={this.handleChanges} ></input>
              <label htmlFor="newcustomer" className="form-label">Nombre</label>
              <input type="text" className="form-control" name="contact_name" id="contact_name" onChange={this.handleChanges}></input>
              <label htmlFor="newcustomer" className="form-label">Cargo</label>
              <input type="text" className="form-control" name="contact_title" id="contact_title" onChange={this.handleChanges}></input>
              <label htmlFor="newcustomer" className="form-label">Dirección</label>
              <input type="text" className="form-control" name="address" id="address" onChange={this.handleChanges}></input>
              <label htmlFor="newcustomer" className="form-label">Ciudad</label>
              <input type="text" className="form-control" name="city" id="city" onChange={this.handleChanges}></input>
              <label htmlFor="newcustomer" className="form-label">Región</label>
              <input type="text" className="form-control" name="region" id="region" onChange={this.handleChanges}></input>
              <label htmlFor="newcustomer" className="form-label">Código Postal</label>
              <input type="text" className="form-control" name="postal_code" id="postal_code" onChange={this.handleChanges}></input>
              <label htmlFor="newcustomer" className="form-label">País</label>
              <input type="text" className="form-control" name="country" id="country" onChange={this.handleChanges}></input>
              <label htmlFor="newcustomer" className="form-label">Teléfono</label>
              <input type="text" className="form-control" name="phone" id="phone" onChange={this.handleChanges}></input>
              <label htmlFor="newcustomer" className="form-label">ID</label>
              <input type="text" className="form-control" name="customer_id" id="customer_id" placeholder="ABCDE" onChange={this.handleChanges}></input>
              
              <button type="button" className="btn btn-primary" onClick={this.reqPost}>Agregar cliente</button>          
              <button type="button" className="btn btn-danger" onClick={this.reqDelete}>Eliminar cliente</button>
              
          </form>
        </div>
      )
    }
  }
  
  export default Customers;
  