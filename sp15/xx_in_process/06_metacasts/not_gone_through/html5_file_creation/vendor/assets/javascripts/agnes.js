// http://www.secretgeek.net/agnes11.asp

var agnes;
(function() {
  var handbag = function(purse) {
    var _agnes = this;
    var _p = {};
    var _myUsualPurse = {
      trimFields : true,
      fieldDelim : ",",
      rowDelim : "\n",
      stripQuote : true,
      qual : '"'
    };
    _p.purse = purse || _myUsualPurse;
    _p.trinkets = {
        trim:function(val) {
            if (val === null) {
                return "";
            }
            return val.toString().replace(/^\s+|\s+$/, "");
        },
        escapeCsv:function(rawText, rowDelim, fieldDelim, qual) {
            //if it contains a fieldDelimeter, a rowDelimiter or a qualfier
            //then it needs to be qualified.
            //(and firstly: any embedded qualifiers need to be doubled)
            var needsQualifying = false;
            if (rawText == null || rawText == undefined) return "";

            if (!rawText.indexOf) { // typeof rawText === "number"
                //maybe a number! no need to escape it.
                return rawText;
            }

            if (rawText.indexOf(fieldDelim) != -1) {
                needsQualifying = true;
            }

            if (rawText.indexOf(rowDelim) != -1) {
                needsQualifying = true;
            }

            if (rawText.indexOf(qual) != -1) {
                needsQualifying = true;
                rawText = rawText.replace(new RegExp(qual, 'g'), qual + qual);
            }

            if (needsQualifying) {
                rawText = qual + rawText + qual;
            }
            return rawText;
        },
        rtrim:function (val, of) {
                if (!of) {
                    of = '\s';
                }
                //if it ends with spaces... remove them!
                of = of.replace('|','\\|');
                of = of.replace('+','\\+');
                //alert(of);
                return val.replace(new RegExp(of + '+$', 'gm'), '');
        }
    }

    _agnes.settings = function(settings) {
      _p.purse = settings || _p.purse;
      return _p.purse;
    };
    _agnes.csvToJson = function(csvString, rowDelim, fieldDelim) {
        var array2d = this.csvToArray2d(csvString, rowDelim, fieldDelim);
        return this.array2dToJson(array2d);
    },

    _agnes.jsonToCsv = function(allObjects, rowDelim, fieldDelim, qual) {
        var properties, result, arr, headings, items;
        rowDelim = rowDelim || this.rowDelimiter();
        fieldDelim = fieldDelim || ',';
        qual = qual || '"';
        properties = [];
        result = "";
        headings = "";
        //if it's not an array, convert it to one now.
        //caveat: what happens if it's a string? (i.e. still has length property)
        if (!allObjects.length) {
          arr = [];
          arr[0] = allObjects;
          allObjects = arr;
        }

        for (var i in allObjects) {
            items = allObjects[i];
            for (var prop in items) {
                if (i == 0)
                {
                    headings += _p.trinkets.escapeCsv(prop, rowDelim, fieldDelim, qual) + fieldDelim;
                }

                result+=_p.trinkets.escapeCsv(items[prop], rowDelim, fieldDelim, qual);
                result+=fieldDelim;
            }

            result = _p.trinkets.rtrim(result, fieldDelim);
            result += rowDelim;
        }
        return _p.trinkets.rtrim(headings, fieldDelim) + rowDelim + result;
    },
    _agnes.csvToArray2d = function(lines, rowDelim, fieldDelim) {
        var lineArray;
        rowDelim = rowDelim || this.rowDelimiter();
        fieldDelim =  fieldDelim || ",";
        lineArray = this.toArray(lines, rowDelim);
        var i, numLines, array2d;
        numLines = lineArray.length;
        array2d = [];
        j = 0;
        for (i = 0; i < numLines; i = i + 1) {
            // skip empty rows.
            if (lineArray[i] == null) { continue;}
            array2d[j] = this.toArray(lineArray[i], fieldDelim);
            j++;
        }

        return array2d;
    },
    _agnes.array2dToJson = function(array2d) {
        var fields, i, numRows, header, jarray, json;
        numRows = array2d.length;
        jarray = [];
        for (i = 0; i < numRows; i = i + 1) {
            fields = (array2d[i]);
            if (i === 0) {
                header=fields;
            } else {
                json = this.rowToJson(header, fields);
                jarray[i-1] = json;
            }
        }
        return jarray;
    },
    _agnes.toArray = function(line, splitOn, unclosedQuoteObject) {
        var unclosedQuote, parts, i, iStart, linei, lineQual;

        if (unclosedQuoteObject != undefined) {
            unclosedQuote = (unclosedQuoteObject.unclosed || false);
        } else {
            unclosedQuote = false;
        }

        var qual = _p.purse.qual || _myUsualPurse.qual;
        var stripQuote = _p.purse.stripQuote || _myUsualPurse.stripQuote;
        splitOn = splitOn || _p.purse.fieldDelim || _myUsualPurse.fieldDelim;
        var trimFields = _p.purse.trimFields || _myUsualPurse.trimFields;
        parts = [];
        i = 0;
        iStart = 0;

        if (line == undefined) {
            return parts;
        }

        for (i = 0; i < line.length; i = i + 1) {
            linei = line.substr(i, splitOn.length);

            lineQual = line.substr(i, qual.length);
            if (linei == splitOn) {
                if (!unclosedQuote) {
                    if (trimFields) {
                        var trimmed = _p.trinkets.trim(line.substr(iStart, i - iStart));
                        parts.push(this.unescapeCsv(trimmed, qual, stripQuote));
                    } else {
                        parts.push(this.unescapeCsv(line.substr(iStart, i - iStart), qual, stripQuote));
                    }
                        iStart = i + splitOn.length;
                }
            } else if (lineQual == qual) {
                unclosedQuote = !unclosedQuote;
                //Maybe... if qual.Length > 1 then
                //iStart = i - 1 + qual.Length;
            }
        }

        if (trimFields) {
            parts.push(this.unescapeCsv(_p.trinkets.trim(line.substr(iStart, i - iStart)), qual, stripQuote));
        } else {
            parts.push(this.unescapeCsv(line.substr(iStart, i - iStart), qual, stripQuote));
        }

        if (unclosedQuoteObject != undefined) {
            unclosedQuoteObject.unclosed = unclosedQuote;
        }

        return parts;
    },
    _agnes.rowDelimiter = function() {
        //determine the standard row delimiter used by this browser
        inputElem = document.createElement('textarea');
        //assign a \r\n to the textarea...
        inputElem.value = "\r\n";
        //then see what value comes back...
        return inputElem.value;
        //strange fact that the browsers silently convert the line endings of a textarea into
        //the line ending of their choice.
        //So far I've only confirmed that chrome and ie do this.
    },
    _agnes.unescapeCsv = function(value, qual, stripQuote) {
            if (value === undefined || value === "") {
                return null;
            }
            if (!stripQuote) {
                return value;
            }
            qual = (qual || '"');
            //Remove leading and trailing whitespace. i.e. spaces and tabs
            if (this.trimFields) {
                value = _p.trinkets.trim(value);
            }
            //If it begins AND ends with quotes -- remove them.
            if (value.substr(0, qual.length) === qual && value.substr(value.length - qual.length) === qual) {
                value = value.substr(qual.length);
                value = value.substr(0, value.length - qual.length);
            }
            //turn double quotes into single quotes
            value = value.replace(new RegExp(qual + qual, 'g'), qual);
            return value;
    },
    _agnes.arrayToJson = function(dataRows) {
        var fields, i, numRows, header, jarray, json;
        numRows = dataRows.length;
        jarray = [];
        for (i = 0; i < numRows; i = i + 1) {
            fields = this.toArray(dataRows[i]);
            if (i === 0) {
                header=fields;
            } else {
                json = this.rowToJson(header, fields);
                jarray[i-1] = json;
            }
        }
        return jarray;
    },
    _agnes.rowToJson = function(properties, values) {
        var i,j,o;//, result;
        o={};
        for (j = 0; j < properties.length; j = j + 1) {
            o[properties[j]]=values[j];
        }
        return o;
    }
  }
  //creates a default handbag for agnes
  /*
  var _defaultHandbag = function() {
    return {
      trimFields : true,
      fieldDelim : ",",
      rowDelim : "\n",
      stripQuote : true,
      qual : '"'
    };
  };
  */
  agnes = new handbag();//_defaultHandbag());
})();