import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { 
  GetFormularios,
  InsertFormulario,
  DeleteFormulario,
  UpdateFormulario,
 } from './componentes/api';

const defaultState = {
  numIdentificacion: "",
  nombre: "",
  apellidos: "",
  tipoIdentificacion:"",
  numBeneficiarios:"",
  //estadoCivil:"",
};


function App() {
  const [data, setData] = useState(defaultState);
  const [listData, setListData] = useState([]);
  const [idx, setIdx] = useState(0);
  const [action, setAction]= useState(0); // 0:none, 1:insert, 2:update
  
  useEffect(() => {
    handlerList(idx);
  }, []);

  const handlerList = async(_idx) => {
    // Realiza una solicitud asincrónica para obtener los formularios
      const response = await GetFormularios();
      // Extrae la propiedad 'result' del objeto de respuesta
      console.log(response);
      const { result } = response;
      console.log(result);
      
      // Extrae la propiedad 'formularios' del objeto 'result'
      const { formularios } = result;
      // console.log({_idx,formularios, n:formularios.length});
      // Si no hay formularios, no hagas nada
      if (formularios.length == 0) {
        //console.log(formularios.length);
        return;
      }
      
      // Si el índice proporcionado es mayor o igual al número de formularios,
      // no hagas nada
      if (_idx >= formularios.length) {
        return;
      }
      
       // Obtiene el formulario correspondiente al índice proporcionado
      const formulario = formularios[_idx];
      //console.log(_idx);
       // Actualiza el estado 'listData' con la lista completa de formularios
      setListData(formularios);
      //console.log(formularios);
      // Actualiza el estado 'data' con el formulario seleccionado
      setData(formulario);
      // Actualiza el estado 'idx' con el índice seleccionado
      setIdx(_idx);
      //console.log(formulario);
      //console.log("Listar");
  }

  const handlerChange = (e) => {
    // Extrae el nombre y el valor del elemento que desencadenó el evento
    const {name, value} = e.target;
    //console.log(data);
     // Si la acción actual es 0 (ninguna), establece la acción en 2 (actualización)
    if ( action == 0) {
      setAction(2);
    }
    // Actualiza el estado 'data' del componente con el nuevo valor del campo modificado,
    // conservando los valores anteriores de 'data' y solo actualizando el campo modificado    
    setData({
      ...data,
      [name]: value,
    });
  };


  const handlerNew = () => {
     // Establece la acción en 1 (insertar)
    setAction(1);
    // Establece los datos del formulario en el estado inicial predeterminado
    setData(defaultState);
    console.log("Insertar");
  };

  const handlerInsert = async () => {
    // Extrae los datos del formulario del estado 'data'
    const { nombre, apellidos, tipoIdentificacion, numBeneficiarios } = data;
    console.log(data);
    //----------------------------------------------------------------------------
    const numBeneficiariosInteger = parseInt(numBeneficiarios);
    console.log(typeof numBeneficiarios);
    //const estadoCivilInteger = parseInt(estadoCivil);
    //console.log(typeof estadoCivil);
    // ---------------------------------------------------------------------------
    // Define valores predeterminados para otros campos del formulario
    //const tipoIdentificacion = "1"
    const estadoCivil = "1";
    const fechaNacimiento = "10/10/2021";
    //const numBeneficiarios = "1";
    const fechaIngreso = "10/10/2023";
    const response = await InsertFormulario(
      nombre,
      apellidos,
      tipoIdentificacion,
      estadoCivil,
      fechaNacimiento,
      numBeneficiarios,
      fechaIngreso,
    );

    const { ok, result } = response;
      if (ok) {
        const _idx = listData.length;
        console.log(result);
        console.log("prueba");
        //console.log({_idx, listData});
        await handlerList(_idx);
        return;
    }
  };

  const handlerUpdate = async () => {
    const { numIdentificacion, nombre, apellidos, tipoIdentificacion, numBeneficiarios }= data;
    const numBeneficiariosInteger = parseInt(numBeneficiarios);
    //const estadoCivilInteger = parseInt(estadoCivil);
    //console.log(typeof estadoCivil);
    const estadoCivil = "1";
    const fechaNacimiento = "10/10/2021";
    //const numBeneficiarios = "1";
    const fechaIngreso = "10/10/2023";
    //const numBeneficiariosEntero = parseInt(numBeneficiarios, 10);
    //const tipoIdentificacion = "CC"
    
    const response = await UpdateFormulario(
      numIdentificacion,
      nombre,
      apellidos,
      tipoIdentificacion,
      estadoCivil,
      fechaNacimiento,
      numBeneficiarios,
      fechaIngreso,
    )

    const { ok, result } = response;
    if (!ok ) {
      console.log(result);
      return;
    }
  };

  const handlerSave = async() => {
    if (action == 1) {
      await handlerInsert();
      setAction(0);
      return;
    } 
    
    if (action==2) {
      await handlerUpdate();
      setAction(0);
      return;
    }
  };


  const handlerDelete = async() => {
    const { numIdentificacion } = data;
    console.log(data);
        
    const response = await DeleteFormulario(numIdentificacion);
    //console.log(numIdentificacion);
    const { ok } = response;
    if (ok ) {
      
      const _idx = listData.length-2;
      await handlerList(_idx);
      return;
    }

    console.log("Eliminar");
  };

    const handlerNext = () => {
      const n = listData.length -1;
      if (idx === n) {
        return;
      }
      const _idx = idx + 1;
      //console.log(n, _idx);
      const _data = listData[_idx];
      setData(_data);
      setIdx(_idx);
  }

  const handlerPrev = () => {
    const n = listData.length;
    if (idx === 0) {
      return;
    }
    const _idx = idx - 1;
    const _data = listData[_idx];
    setData(_data);
    setIdx(_idx);
  }

// <option selected>Open this select menu</option>
// <option defaultValue=" ">Open this select menu</option>

  return (
    <>
      <div className='container'>
        <div className="card mt-5">
          <div className="card-body">
            <h5 className="card-title">Formulario</h5>
            <div className="mb-3">
              <label className="form-label">Numero de Identificacion</label>
              <div className="input-group">
               <input
                type="text"
                className="form-control"
                id="basic-url"
                aria-describedby="basic-addon3 basic-addon4"
                name="numIdentificacion"
                value = {data.numIdentificacion}
                onChange = {(e)=> handlerChange(e)}/>
              </div>
            </div>
          <div className="mb-3">
            <label className="form-label">Nombres</label>
            <div className="input-group">
               <input
                type="text"
                className="form-control"
                id="basic-url"
                aria-describedby="basic-addon3 basic-addon4"
                name="nombre"
                value = {data.nombre}
                onChange = {(e)=> handlerChange(e)}/>
            </div>
          </div>
          <div className="mb-3">
            <label className="form-label">Apellidos</label>
            <div className="input-group">
               <input
                type="text"
                className="form-control"
                id="basic-url"
                aria-describedby="basic-addon3 basic-addon4"
                name="apellidos"
                value = {data.apellidos}
                onChange = {(e)=> handlerChange(e)}
              />
            </div>
          </div>
          <div className="mb-3">
            <label className="form-label">Estado Civil</label>
            <div className="input-group">
              <select className="form-select" aria-label="Default select example"
                  name="estadoCivil"
                  //defaultValue="open"
                  value = {data.estadoCivil}
                  onChange = {(e)=> handlerChange(e)}>
                <option value="open">Open this select menu</option>
                <option value="1">Soltero</option>
                <option value="2">Casado</option>
                <option value="3">Union Libre</option>
              </select>
            </div>
          </div>
          
          <div className="mb-3">
            <label className="form-label">Nro. Beneficiarios</label>
            <div className="input-group">
               <input
                //type="number"
                type="text"
                className="form-control"
                id="basic-url"
                aria-describedby="basic-addon3 basic-addon4"
                name="numBeneficiarios"
                value = {data.numBeneficiarios}
                onChange = {(e)=> handlerChange(e)}
              />
           </div>
            
          </div>
          <div className="mb-3">
            <label className="form-label">Tipo de identificacion</label>
            <div className="input-group">
              <select className="form-select" aria-label="Default select example"
                  name="tipoIdentificacion"
                  //defaultValue="open"
                  value = {data.tipoIdentificacion}
                  onChange = {(e)=> handlerChange(e)}>
                <option value="open">Open this select menu</option>
                <option value="1">Tarjeta de identidad</option>
                <option value="2">Cedula de ciudadania</option>
                <option value="3">Pasaporte</option>
              </select>
            </div>
          </div>
        </div>
      <div className="card-body">
        <button
          type="button"
          disabled=""
          className="celsia-primary celsia-btn celsia-round"
          onClick={() => handlerNew()}
        >Nuevo</button>

        <button
          type="button"
          disabled=""
          className={ action != 0 ? "celsia-primary celsia-btn celsia-round ": "celsia-primary celsia-btn celsia-round visually-hidden" }
          onClick={() => handlerSave()}
        >Guardar</button>

        <button
          type="button"
          disabled=""
          className="btn btn-danger celsia-btn celsia-round"
          onClick={() => handlerDelete()}
        >Eliminar</button>

        <button
          type="button"
          disabled=""
          className="btn btn-danger celsia-btn celsia-round"
          onClick={() => handlerNext()}
        >Siguiente</button>

        <button
          type="button"
          disabled=""
          className="btn btn-danger celsia-btn celsia-round"
          onClick={() => handlerPrev()}
        >Anterior
        </button>
      </div>
    </div>
    </div>
      
    </>
  );
};

export default App;
