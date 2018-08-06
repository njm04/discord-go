package main

import (
	"github.com/james-bowman/nlp"
	"gonum.org/v1/gonum/mat"
	"fmt"
	"github.com/james-bowman/nlp/measures/pairwise"
)

func main() {
	testCorpus := []string{
		"The quick brown fox jumped over the lazy dog",
		"hey diddle diddle, the cat and the fiddle",
		"the fast cunning brown fox liked the slow canine dog ",
		"the little dog laughed to see such fun",
		"and the dish ran away with the spoon",
	}

	query := "the cunning creature ran around the canine"

	vectoriser := nlp.NewCountVectoriser()
	transformer := nlp.NewTfidfTransformer()

	reducer := nlp.NewTruncatedSVD(2)

	matrix, _ := vectoriser.FitTransform(testCorpus...)

	queryMat, _ := vectoriser.Transform(query)
	calcCosine(queryMat, matrix, testCorpus, "Raw TF")

	tfidfmat, _ := transformer.FitTransform(matrix)
	tfidfquery, _ := transformer.Transform(queryMat)
	calcCosine(tfidfquery, tfidfmat, testCorpus, "TF-IDF")

	lsi, _ := reducer.FitTransform(tfidfmat)
	queryVector, _ := reducer.Transform(tfidfquery)
	calcCosine(queryVector, lsi, testCorpus, "LSA")
}

func calcCosine(query mat.Matrix, tdmat mat.Matrix, corpus []string, name string) {
	_, docs := tdmat.Dims()
	var matched int
	highestSimilarity := -1.0

	fmt.Printf("Comparing based on %s\n", name)

	for i := 0; i< docs; i++ {
		queryVec := query.(mat.ColViewer).ColView(0)
		docVec := tdmat.(mat.ColViewer).ColView(i)
		similarity := pairwise.CosineSimilarity(queryVec, docVec)
		fmt.Printf("Comparing '%s' = %f\n", corpus[i], similarity)
		if similarity > highestSimilarity {
			matched = i
			highestSimilarity = similarity
		}
	}
	fmt.Printf("Matched '%s'", corpus[matched])
}