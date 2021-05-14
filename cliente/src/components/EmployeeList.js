import React from "react";

const Employees = ({employees}) =>{




    return(

        <table className="table">
        <thead>
          <tr >
            <th scope="col">Empleado ID</th>
            <th scope="col">Apellidos</th>
            <th scope="col">Nombre</th>
            <th scope="col">Cargo</th>
            <th scope="col">Fecha de contratación</th>
            <th scope="col">Dirección</th>
            <th scope="col">Ciudad</th>
            <th scope="col">Región</th>
            <th scope="col">Código Postal</th>
            <th scope="col">País</th>
            <th scope="col">Teléfono</th>
          </tr>
        </thead>
        <tbody>
            {employees.map((employee)=>(
              
              <tr key={employee.employee_id}>
              <th scope="row">{employee.employee_id}</th>
              <td>{employee.first_name}</td>
              <td>{employee.last_name}</td>
              <td>{employee.title}</td>
              <td>{employee.hire_date}</td>
              <td>{employee.address}</td>
              <td>{employee.city}</td>
              <td>{employee.region}</td>
              <td>{employee.postal_code}</td>
              <td>{employee.country}</td>
              <td>{employee.phone}</td>
            </tr>
           
            ))}
        </tbody>
  </table>
    )
};

export default Employees

  


  
