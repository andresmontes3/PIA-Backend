import React from "react";
import ProductList from '../components/ProductList';


class Products extends React.Component{

    state = {
      products: [],
        product_id: 0,
        product_name: '',
        quantity_per_unit: '',
        unit_price: 0.0,
        units_in_stock: 0,
        discontinued: false,
     };
  
     handleChanges = async(e)=>{
       const target = e.target
        await this.setState({
         [target.name]:target.value
       });
     }
  
  
    
     reqDelete=()=>{
      fetch('http://localhost:1323/products?product_id='+this.state.product_id,{
        method:'delete'
     })
     .then((data) => {
      this.reqGet()
       console.log(data)
     })
     .catch(console.log)
     }
  
     reqGet=()=>{
      fetch('http://localhost:1323/products',{
        method:'get'
     })
     .then(res => res.json())
     .then((data) => {
       console.log(data)
       this.setState({ products: data })
     })
     .catch(console.log)
     }
  
     reqPost=async()=>{
      await fetch('http://localhost:1323/products',{
        method:'post',
        headers: {
          'Accept': 'application/json, text/plain, */*',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(
          {
            product_id : parseInt(this.state.product_id),
            product_name : this.state.product_name,
            quantity_per_unit : this.state.quantity_per_unit,
            unit_price : parseFloat(this.state.unit_price),
            units_in_stock : parseInt(this.state.units_in_stock),
            discontinued : this.state.discontinued === "true" ? true:false, 
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
        <h2>Lista Productos</h2>
        <ProductList products={this.state.products}/>
        <h3>Información Nuevo Producto</h3>
        <form>   
              <label htmlFor="newproduct" className="form-label">Nombre</label>
              <input type="text" className="form-control" name="product_name" id="product_name" onChange={this.handleChanges} ></input>
              <label htmlFor="newproduct" className="form-label">Cantidad por unidad</label>
              <input type="text" className="form-control" name="quantity_per_unit" id="quantity_per_unit" onChange={this.handleChanges}></input>
              <label htmlFor="newproduct" className="form-label">Precio por unidad</label>
              <input type="text" className="form-control" name="unit_price" id="unit_price" onChange={this.handleChanges}></input>
              <label htmlFor="newproduct" className="form-label">Unidades en stock</label>
              <input type="text" className="form-control" name="units_in_stock" id="units_in_stock" onChange={this.handleChanges}></input>
              <label htmlFor="newproduct" className="form-label" placeholder="true o false">Descontinuado</label>
              <input type="text" className="form-control" name="discontinued" id="discontinued" onChange={this.handleChanges}></input>
              <label htmlFor="newproduct" className="form-label">ID</label>
              <input type="text" className="form-control" name="product_id" id="product_id" placeholder="12345" onChange={this.handleChanges}></input>
              
              <button type="button" className="btn btn-primary" onClick={this.reqPost}>Agregar producto</button>          
              <button type="button" className="btn btn-danger" onClick={this.reqDelete}>Eliminar producto</button>
              
          </form>
        </div>
      )
    }
  }
  
  export default Products;
  