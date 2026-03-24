- It'll be called once -> useEffect( () => {

}, [])

- But if there is some parameters inside [], then whenever the parameter value changes then it'll be called again and again

- Dependencies of useEffect -> are the parameters inside [] 