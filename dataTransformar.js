const productSales = require('./Data/ProductSales.json');
const { groupBy, transformDataToObj, sum } = require('./utils');

const mapFunctions = {
    'sum': sum
}

function doSummation(products, measures) {
    Object.keys(products).forEach((key) => {
        if (Array.isArray(products[key])) {
            const transformedDataResult = {};
            measures.forEach((measure) => {
                const total = mapFunctions[measure.fun](products[key], measure.name);
                transformedDataResult[measure.name] = total;
            });
            products[key] = transformedDataResult;
        } else {
            doSummation(products[key], measures);
        }
    });
    return products;
}

function doTransform(dimension, products) {
    Object.keys(products).forEach((key) => {
        if (Array.isArray(products[key])) {
            if (key === dimension) {
                products[key] = groupBy(products[key], dimension);
            } else {
                products[key] = { [dimension]: groupBy(products[key], dimension) }
            }
        } else {
            doTransform(dimension, products[key]);
        }
    });
    return products;
}

function dataTransformar() {
    const dimensions = ['category', 'rating'];
    let result = { [dimensions[0]]: productSales }
    dimensions.forEach((dimension) => {
        result = doTransform(dimension, result)
    });

    const groupByColumns = ['year', 'month'];
    groupByColumns.forEach((dimension) => {
        result = doTransform(dimension, result)
    });

    const measures = [{ name: 'revenue', fun: 'sum' }, { name: 'tax', fun: 'sum' }];
    result = doSummation(result, measures)
    return result;
}

console.log(JSON.stringify(dataTransformar(), null, 2));