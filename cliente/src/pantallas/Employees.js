import React from "react";
import EmployeeList from '../components/EmployeeList';


class Employees extends React.Component{

    state = {
      token:'',
      employees: [],
        employee_id: 0,
        last_name: '',
        first_name: '',
        title: '',
        hire_date: '',
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
      fetch('http://localhost:1323/api/employees?employee_id='+this.state.employee_id,{
        method:'delete',
        headers:{
          'Authorization':'Bearer '+this.state.token
        }
     })
     .then((data) => {
      this.reqGet()
       console.log(data)
     })
     .catch(console.log)
     }
  
     reqGet=()=>{
      fetch('http://localhost:1323/api/employees',{
        method:'get',
        headers:{
          'Authorization':'Bearer '+this.state.token
        }
     })
     .then(res => res.json())
     .then((data) => {
       console.log(data)
       this.setState({ employees: data })
     })
     .catch(console.log)
     }
  
     reqPost=async()=>{
      await fetch('http://localhost:1323/api/employees',{
        method:'post',
        headers: {
          'Accept': 'application/json, text/plain, */*',
          'Content-Type': 'application/json',
          'Authorization':'Bearer '+this.state.token
        },
        body: JSON.stringify(
          {
            employee_id : parseInt(this.state.employee_id),
            last_name : this.state.last_name,
            first_name : this.state.first_name,
            title : this.state.title,
            hire_date : this.state.hire_date,
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


     login = async()=>{
      await fetch('http://localhost:1323/login?username=josuefdz&password=contra',{
        method:'post'
     })
     .then((res) => {

       res.json().then(result =>{
         console.log(result.token)
         this.setState({token:result.token})
          this.reqGet()
       })
       
     })
     .catch(console.log)
     }
   
     componentDidMount() {
       this.login();
     }
     
  
     render(){
  
      return (
  
  
        <div className="container">
        <h2>Lista Empleados</h2>
        <EmployeeList employees={this.state.employees}/>
        <h3>Información Nuevo Empleado</h3>
        <form>   
              <label htmlFor="newemployee" className="form-label">Apellidos</label>
              <input type="text" className="form-control" name="last_name" id="first_name" onChange={this.handleChanges} ></input>
              <label htmlFor="newemployee" className="form-label">Nombre</label>
              <input type="text" className="form-control" name="first_name" id="last_name" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">Cargo</label>
              <input type="text" className="form-control" name="title" id="title" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">Fecha de contratación</label>
              <input type="text" className="form-control" name="hire_date" id="hire_date" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">Dirección</label>
              <input type="text" className="form-control" name="address" id="address" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">Ciudad</label>
              <input type="text" className="form-control" name="city" id="city" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">Región</label>
              <input type="text" className="form-control" name="region" id="region" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">Código Postal</label>
              <input type="text" className="form-control" name="postal_code" id="postal_code" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">País</label>
              <input type="text" className="form-control" name="country" id="country" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">Teléfono</label>
              <input type="text" className="form-control" name="phone" id="phone" onChange={this.handleChanges}></input>
              <label htmlFor="newemployee" className="form-label">ID</label>
              <input type="text" className="form-control" name="employee_id" id="employee_id" placeholder="12345" onChange={this.handleChanges}></input>
              
              <button type="button" className="btn btn-primary" onClick={this.reqPost}>Agregar empleado</button>          
              <button type="button" className="btn btn-danger" onClick={this.reqDelete}>Eliminar empleado</button>
              
          </form>
        </div>
      )
    }
  }
  
  export default Employees;
  