const productSales = require('./Data/ProductSales.json');
const { groupBy, sum } = require('./utils');

const mapFunctions = {
    'sum': sum
}

function aggregateData(products, key, options) {
    const transformedDataResult = {};
    options.measures.forEach((measure) => {
        const total = mapFunctions[measure.fun](products[key], measure.name);
        transformedDataResult[measure.name] = total;
    });
    return transformedDataResult;
}

function groupData(products, key, options) {
    const column = options.column;
    if (key === column) {
        return groupBy(products[key], column);
    } else {
        return { [column]: groupBy(products[key], column) }
    }
}

function findArrayRecursivelyandDoCallback(products, options, cb) {
    Object.keys(products).forEach((key) => {
        if (Array.isArray(products[key])) {
            products[key] = cb(products, key, options);
        } else {
            findArrayRecursivelyandDoCallback(products[key], options, cb);
        }
    });
    return products;
}

function dataTransformarTwoDimension(productSales) {
    const dimensions = ['category', 'rating'];
    let result = { [dimensions[0]]: productSales }
    dimensions.forEach((dimension) => {
        const options = { column: dimension };
        result = findArrayRecursivelyandDoCallback(result, options, groupData);
    });

    const groupByColumns = ['year', 'month'];
    groupByColumns.forEach((dimension) => {
        const options = { column: dimension };
        result = findArrayRecursivelyandDoCallback(result, options, groupData);
    });

    const measures = [{ name: 'revenue', fun: 'sum' }, { name: 'tax', fun: 'sum' }];
    const options = { measures: measures };
    result = findArrayRecursivelyandDoCallback(result, options, aggregateData)
    return result;
}

function dataTransformarOneDimension(productSales) {
    const dimensions = ['category'];
    let result = { [dimensions[0]]: productSales }
    dimensions.forEach((dimension) => {
        const options = { column: dimension };
        result = findArrayRecursivelyandDoCallback(result, options, groupData);
    });

    const groupByColumns = ['month'];
    groupByColumns.forEach((dimension) => {
        const options = { column: dimension };
        result = findArrayRecursivelyandDoCallback(result, options, groupData);
    });

    const measures = [{ name: 'revenue', fun: 'sum' }, { name: 'tax', fun: 'sum' }];
    const options = { measures: measures };
    result = findArrayRecursivelyandDoCallback(result, options, aggregateData)
    return result;
}
console.log("***********************One Dimenstion*******************");
console.log(JSON.stringify(dataTransformarOneDimension(productSales), null, 2));
console.log("\n\n***********************Two Dimenstion*******************");
console.log(JSON.stringify(dataTransformarTwoDimension(productSales), null, 2));