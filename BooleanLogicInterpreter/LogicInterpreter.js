const interpreter = (expression) => {
    let variables = {};
  
    const isVariable = (str) => {
      return variables[str] !== undefined;
    };
  
    const evaluate = (str) => {
      if (str === 'T') {
        return true;
      }
      if (str === 'F') {
        return false;
      }
      if (isVariable(str)) {
        return variables[str];
      }
      return str;
    };
  
    const not = (value) => {
      return !value;
    };
  
    const and = (left, right) => {
      return left && right;
    };
  
    const or = (left, right) => {
      return left || right;
    };
  
    const evaluateExpression = (expression) => {
      if (expression.includes('(')) {
        let openIndex = expression.indexOf('(');
        let closeIndex = expression.indexOf(')');
        let subExpression = expression.substring(openIndex + 1, closeIndex);
        expression = expression.substring(0, openIndex) + evaluateExpression(subExpression) + expression.substring(closeIndex + 1);
      }
  
      if (expression.includes('¬')) {
        let notIndex = expression.indexOf('¬');
        let value = evaluate(expression.substring(notIndex + 1));
        return not(value);
      }
  
      if (expression.includes('∧')) {
        let andIndex = expression.indexOf('∧');
        let left = evaluate(expression.substring(0, andIndex));
        let right = evaluate(expression.substring(andIndex + 1));
        return and(left, right);
      }
  
      if (expression.includes('∨')) {
        let orIndex = expression.indexOf('∨');
        let left = evaluate(expression.substring(0, orIndex));
        let right = evaluate(expression.substring(orIndex + 1));
        return or(left, right);
      }
  
      return evaluate(expression);
    };
  
    const assignVariable = (expression) => {
      let equalIndex = expression.indexOf('=');
      let variable = expression.substring(4, equalIndex).trim();
      let value = evaluateExpression(expression.substring(equalIndex + 1));
      variables[variable] = value;
      return variables[variable];
    };
  
    if (expression.startsWith('let ')) {
      return assignVariable(expression);
    }
    return evaluateExpression(expression);
  };
  
  console.log(interpreter('T ∨ F'));
  console.log(interpreter('T ∧ F'));
  console.log(interpreter('(T ∧ F) = F'));
  console.log(interpreter('let X = F'));
  console.log(interpreter('let Y = ¬X'));
  console.log(interpreter('¬X ∧ Y'));