const Select = (props) => {
    return (
        <div className="mb-3">
        <label htmlFor={props.name} className='form-label'>
            {" "}
            { props.title}{" "}
            </label>
        <select name={props.name} 
                className="form-select" 
                value={props.name} 
                onChange={props.handleChange}
                >

            <option className='form-select' value="">{props.placeholder}</option>
            {props.options.map((option) => {
                return (
                    <option 
                    className='form-select' 
                    value={option.id}
                    key={option.id}
                    label={option.value}
                    >
                        {option.value}
                    </option>
                );
            })}
        </select>
    </div>
    )
}

export default Select;