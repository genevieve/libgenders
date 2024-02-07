package libgenders_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ryanmoran/libgenders"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDatabase(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
		path   string
	)

	it.Before(func() {
		file, err := os.CreateTemp("", "genders")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		path = file.Name()
	})

	it.After(func() {
		Expect(os.Remove(path)).To(Succeed())
	})

	context("NewDatabase", func() {
		it("loads the database from a path on disk", func() {
			_, err := libgenders.NewDatabase(path)
			Expect(err).NotTo(HaveOccurred())
		})

		context("failure cases", func() {
			context("when the filepath does not exist", func() {
				it("returns an error", func() {
					_, err := libgenders.NewDatabase("no-such-file")
					Expect(err).To(MatchError(ContainSubstring("no-such-file: no such file or directory")))
				})
			})
		})
	})

	context("GetNodes", func() {
		var (
			testdata = []string{
				"genders.base",
				"genders.base_comma",
				"genders.base_comments_beginning_of_line",
				"genders.base_comments_beginning_of_line_comma",
				"genders.base_comments_beginning_of_line_hostrange",
				"genders.base_comments_end_of_line",
				"genders.base_comments_end_of_line_comma",
				"genders.base_comments_end_of_line_hostrange",
				"genders.base_comments_end_of_line_with_whitespace",
				"genders.base_comments_end_of_line_with_whitespace_comma",
				"genders.base_comments_end_of_line_with_whitespace_hostrange",
				"genders.base_comments_middle_of_line",
				"genders.base_comments_middle_of_line_comma",
				"genders.base_comments_middle_of_line_hostrange",
				"genders.base_hostrange",
				"genders.base_hostrange_single",
				"genders.base_whitespace_after_attrs",
				"genders.base_whitespace_after_attrs_comma",
				"genders.base_whitespace_after_attrs_hostrange",
				"genders.base_whitespace_after_nodes",
				"genders.base_whitespace_after_nodes_comma",
				"genders.base_whitespace_after_nodes_hostrange",
				"genders.base_whitespace_before_and_after_nodes",
				"genders.base_whitespace_before_and_after_nodes_comma",
				"genders.base_whitespace_before_and_after_nodes_hostrange",
				"genders.base_whitespace_before_nodes",
				"genders.base_whitespace_before_nodes_comma",
				"genders.base_whitespace_before_nodes_hostrange",
				"genders.base_whitespace_between_nodes",
				"genders.base_whitespace_between_nodes_and_attrs",
			}
		)

		for _, filename := range testdata {
			var database libgenders.Database
			path := filepath.Join("./testdata", filename)

			it.Before(func() {
				var err error
				database, err = libgenders.NewDatabase(path)
				Expect(err).NotTo(HaveOccurred())
			})

			context(fmt.Sprintf("given %s", path), func() {
				it("returns a list of nodes", func() {
					nodes := database.GetNodes()
					Expect(nodes).To(Equal([]libgenders.Node{
						{
							Name: "node1",
							Attributes: map[string]string{
								"attr1": "",
								"attr2": "val2",
							},
						},
						{
							Name: "node2",
							Attributes: map[string]string{
								"attr1": "",
								"attr2": "val2",
							},
						},
					}))
				})
			})
		}

		context("when the file only contains node names", func() {
			var database libgenders.Database

			it.Before(func() {
				var err error
				database, err = libgenders.NewDatabase("./testdata/genders.nodes_only_many")
				Expect(err).NotTo(HaveOccurred())
			})

			it("returns a list of nodes", func() {
				nodes := database.GetNodes()
				Expect(nodes).To(Equal([]libgenders.Node{
					{Name: "node1"},
					{Name: "node2"},
					{Name: "node3"},
					{Name: "node4"},
					{Name: "node5"},
				}))
			})
		})

		context("nodes share common attributes", func() {
			var (
				testdata = []string{
					"genders.query_1",
					"genders.query_1_comma",
					"genders.query_1_hostrange",
				}
			)

			for _, filename := range testdata {
				var database libgenders.Database
				path := filepath.Join("./testdata", filename)

				it.Before(func() {
					var err error
					database, err = libgenders.NewDatabase(path)
					Expect(err).NotTo(HaveOccurred())
				})

				context(fmt.Sprintf("given %s", path), func() {
					it("returns a list of nodes", func() {
						nodes := database.GetNodes()
						Expect(nodes).To(Equal([]libgenders.Node{
							{
								Name: "node1",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr3": "",
									"attr4": "val4",
									"attr7": "",
									"attr8": "val8",
								},
							},
							{
								Name: "node2",
								Attributes: map[string]string{
									"attr1":  "",
									"attr2":  "val2",
									"attr3":  "",
									"attr4":  "val4",
									"attr9":  "",
									"attr10": "val10",
								},
							},
							{
								Name: "node3",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr3": "",
									"attr4": "val4",
									"attr7": "",
									"attr8": "val8",
								},
							},
							{
								Name: "node4",
								Attributes: map[string]string{
									"attr1":  "",
									"attr2":  "val2",
									"attr3":  "",
									"attr4":  "val4",
									"attr9":  "",
									"attr10": "val10",
								},
							},
							{
								Name: "node5",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr5": "",
									"attr6": "val6",
									"attr7": "",
									"attr8": "val8",
								},
							},
							{
								Name: "node6",
								Attributes: map[string]string{
									"attr1":  "",
									"attr2":  "val2",
									"attr5":  "",
									"attr6":  "val6",
									"attr9":  "",
									"attr10": "val10",
								},
							},
							{
								Name: "node7",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr5": "",
									"attr6": "val6",
									"attr7": "",
									"attr8": "val8",
								},
							},
							{
								Name: "node8",
								Attributes: map[string]string{
									"attr1":  "",
									"attr2":  "val2",
									"attr5":  "",
									"attr6":  "val6",
									"attr9":  "",
									"attr10": "val10",
								},
							},
						}))
					})
				})
			}
		})
	})
}
