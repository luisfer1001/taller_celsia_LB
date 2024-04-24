export const GetFormularios = async function() {
    const url = "http://localhost:3000/formularios";
    let requestOptions = {
        method: "GET",
        headres: {
            "Content-Type": "application/json",
        },
    };

    const c = await fetch(url, requestOptions)
        .then((response) => {
            const b = response.json().then((a) => {
                //console.log(a);
                return a
            });
            return b;
         })
        .catch((error) => {
            return {
                ok: false,
                result: {
                    message: "Error en la consulta: " + error,
                },
            };
        });

    return c;
};

export const InsertFormulario = async function (
    nombre,
    apellidos,
    tipoIdentificacion,
    estadoCivil,
    fechaNacimiento,
    numBeneficiarios,
    fechaIngreso,
    ){

   
    if (nombre ==="" || apellidos === ""){
        return {
            ok: false,
            result: {
                    message: "El nombre y apellidos son obligatorios",
            },
        };
    }

    const data = {
        nombre,
        apellidos,
        tipoIdentificacion,
        estadoCivil,
        fechaNacimiento,
        numBeneficiarios,
        fechaIngreso,
    };
     console.log(data);


    const url = "http://localhost:3000/formularios";
    let requestOptions = {
        method: "POST",
        headres: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    const c = await fetch(url, requestOptions)
        .then((response) => {
            const b = response.json().then((a) => {
                console.log(a);
                return a
            });
            return b;
         })
        .catch((error) => {
            return {
                ok: false,
                result: {
                    message: "Error en la consulta: " + error,
                },
            };
        });

    return c;

};

export const UpdateFormulario  = async function (
    numIdentificacion,
    nombre,
    apellidos,
    tipoIdentificacion,
    estadoCivil,
    fechaNacimiento,
    numBeneficiarios,
    fechaIngreso, 
) {
    if (nombre ==="" || apellidos === ""){
        return {
            ok: false,
            result: {
                    message: "El nombre y apellidos son obligatorios",
            },
        };
    }

    

    const data = {
        nombre,
        apellidos,
        tipoIdentificacion,
        estadoCivil,
        fechaNacimiento,
        numBeneficiarios,
        fechaIngreso,
    };

    

    const url = "http://localhost:3000/formularios/"+numIdentificacion;
    console.log(url);

    let requestOptions = {
        method: "PUT",
        headres: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    };

    const c = await fetch(url, requestOptions)
        .then((response) => {
            const b = response.json().then((a) => {
                console.log(a);
                return a
            });
            return b;
         })
        .catch((error) => {
            return {
                ok: false,
                result: {
                    message: "Error en la consulta: " + error,
                },
            };
        });

    return c;
}

export const DeleteFormulario = async function (id){
    const url = "http://localhost:3000/formularios/" + id;
    //console.log(id);
    let requestOptions = {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
    };

    const c = await fetch(url, requestOptions)
        .then((response) => {
        const b = response.json().then((a) => {
            console.log(a);
            console.log(url);
            return a
        });
        return b;
     })
    .catch((error) => {
        return {
            ok: false,
            result: {
                message: "Error en la consulta: " + error,
            },
        };
    });
    
    return c;
}

